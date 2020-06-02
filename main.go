package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/patst/alexa-skills-kit-for-go/alexa"
)

var skill alexa.Skill

func main() {
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
		Addr:         "localhost:8192",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Starting webserver")
	log.Fatal(srv.ListenAndServe())
}

func intentDispatchHandler(request *alexa.IntentRequest, response *alexa.ResponseEnvelope) {
	switch request.Intent.Name {
	case "ConspiracyIntent":
		handleConspiracyIntent(request, response)
	case "AMAZON.HelpIntent":
		handleHelpIntent(request, response)
	case "AMAZON.StopIntent":
		handleStopOrCancelIntent(request, response)
	case "AMAZON.CancelIntent":
		handleStopOrCancelIntent(request, response)
	case "AMAZON.YesIntent":
		handleYesIntent(request, response)
	default:
		log.Println("Unknown intent!", request.Intent.Name)
	}
}

const welcomeReprompt = "Du kannst dir die neusten Wahrheiten erzählen lassen. Sage dazu, Sag mir die Wahrheit"
const welcome = "Willkommen beim mächtigen Aluhut. Du kannst dir die absoluten Wahrheiten erzählen lassen. Sage dazu, Sag mir die Wahrheit"
const goodbye = "Auf Wiedersehen und bleib skeptisch!"
const moreFacts = "Möchtest du noch eine Wahrheit hören?"

func handleYesIntent(request *alexa.IntentRequest, responseEnvelope *alexa.ResponseEnvelope) {
	handleConspiracyIntent(request, responseEnvelope)
}

func handleConspiracyIntent(request *alexa.IntentRequest, responseEnvelope *alexa.ResponseEnvelope) {
	// TODO
	randomFact := "TODO"
	moreFacts := moreFacts

	responseEnvelope.Response.SetOutputSpeech(randomFact + " " + moreFacts)
	responseEnvelope.Response.SetReprompt(moreFacts)
}

func handleStopOrCancelIntent(request *alexa.IntentRequest, responseEnvelope *alexa.ResponseEnvelope) {
	responseEnvelope.Response.SetOutputSpeech(goodbye)
}

func handleHelpIntent(request *alexa.IntentRequest, responseEnvelope *alexa.ResponseEnvelope) {
	responseEnvelope.Response.SetOutputSpeech(welcomeReprompt)
	responseEnvelope.Response.SetReprompt(welcomeReprompt)
}

func launchRequestHandler(request *alexa.LaunchRequest, responseEnvelope *alexa.ResponseEnvelope) {

	responseEnvelope.Response.SetOutputSpeech(welcome)
	responseEnvelope.Response.SetReprompt(welcomeReprompt)
}

func sessionEndedRequestHandler(request *alexa.SessionEndedRequest, responseEnvelope *alexa.ResponseEnvelope) {
}
