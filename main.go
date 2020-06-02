package main

import (
	"flag"
	"log"
	"math/rand"
	"mkammerer.de/go-maechtiger-aluhut/facts"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/patst/alexa-skills-kit-for-go/alexa"
)

var skill alexa.Skill

func main() {
	// Seed RNG
	rand.Seed(time.Now().UnixNano())

	address := flag.String("address", "localhost:8192", "address to listen on")
	help := flag.Bool("help", false, "prints this help")
	flag.Parse()
	if *help {
		flag.Usage()
		return
	}

	router := mux.NewRouter()
	skill = alexa.Skill{
		ApplicationID:  "amzn1.ask.skill.789b02dc-3997-4f0d-8182-140aec7d9296",
		OnIntent:       intentDispatchHandler,
		OnLaunch:       launchRequestHandler,
		OnSessionEnded: sessionEndedRequestHandler,
	}

	router.Handle("/maechtiger-aluhut", skill.GetHTTPSkillHandler()).Methods("POST")

	srv := &http.Server{
		Handler:      router,
		Addr:         *address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Running on " + *address)
	log.Fatal(srv.ListenAndServe())
}

func intentDispatchHandler(request *alexa.IntentRequest, response *alexa.ResponseEnvelope) {
	switch request.Intent.Name {
	case "ConspiracyIntent":
		handleConspiracyIntent(request, response)
	case "AMAZON.YesIntent":
		handleConspiracyIntent(request, response)

	case "AMAZON.HelpIntent":
		handleHelpIntent(request, response)

	case "AMAZON.StopIntent":
		handleStopIntent(request, response)
	case "AMAZON.CancelIntent":
		handleStopIntent(request, response)
	case "AMAZON.NoIntent":
		handleStopIntent(request, response)
	default:
		log.Println("Unknown intent!", request.Intent.Name)
	}
}

const welcomeReprompt = "Du kannst dir die neusten Wahrheiten erzählen lassen. Sage dazu, Sag mir die Wahrheit"
const welcome = "Willkommen beim mächtigen Aluhut. Du kannst dir die absoluten Wahrheiten erzählen lassen. Sage dazu, Sag mir die Wahrheit"
const goodbye = "Auf Wiedersehen und bleib skeptisch!"
const moreFacts = "Möchtest du noch eine Wahrheit hören?"

// Used to set a pointer on it
var aTrue = true
var aFalse = false

func handleConspiracyIntent(request *alexa.IntentRequest, responseEnvelope *alexa.ResponseEnvelope) {
	log.Println("handleConspiracyIntent()")

	randomFact := facts.RandomFact()

	responseEnvelope.Response.SetOutputSpeech(randomFact + " " + moreFacts)
	responseEnvelope.Response.SetReprompt(moreFacts)
	responseEnvelope.Response.ShouldEndSession = &aFalse
}

func handleStopIntent(request *alexa.IntentRequest, responseEnvelope *alexa.ResponseEnvelope) {
	log.Println("handleStopIntent()")

	responseEnvelope.Response.SetOutputSpeech(goodbye)
	responseEnvelope.Response.ShouldEndSession = &aTrue
}

func handleHelpIntent(request *alexa.IntentRequest, responseEnvelope *alexa.ResponseEnvelope) {
	log.Println("handleHelpIntent()")

	responseEnvelope.Response.SetOutputSpeech(welcomeReprompt)
	responseEnvelope.Response.SetReprompt(welcomeReprompt)
	responseEnvelope.Response.ShouldEndSession = &aFalse
}

func launchRequestHandler(request *alexa.LaunchRequest, responseEnvelope *alexa.ResponseEnvelope) {
	log.Println("launchRequestHandler()")

	responseEnvelope.Response.SetOutputSpeech(welcome)
	responseEnvelope.Response.SetReprompt(welcomeReprompt)
	responseEnvelope.Response.ShouldEndSession = &aFalse
}

func sessionEndedRequestHandler(request *alexa.SessionEndedRequest, responseEnvelope *alexa.ResponseEnvelope) {
	log.Println("sessionEndedRequestHandler()")
	responseEnvelope.Response.ShouldEndSession = &aTrue
}
