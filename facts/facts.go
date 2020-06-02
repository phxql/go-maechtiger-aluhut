package facts

import (
	"math/rand"
	"strings"
)

const Facts = `Tupac Shakur ist von der US-Regierung getötet worden; alternativ ist sein Tod nur vorgetäuscht.
Das FBI ist für ein Feuer verantwortlich, das sich im Rahmen einer Belagerung einer Branch-Davidians-Anlage entfachte. Durch das Feuer starb eine zweistellige Zahl an Menschen, darunter auch Kinder.
Die Vril-Gesellschaft ist mit übernatürlichen Methoden in der Lage gewesen, den Nationalsozialismus aufleben zu lassen.
Es sank nicht die Titanic, sondern ihr Schwesterschiff Olympic.
Der Unfall, der zum Tod Dianas führte, ist durch den Secret Intelligence Service M ei six herbeigeführt worden.
Die radikale Organisation der Sonnentempler ist durch Rechtsextremismus unterwandert.
Slawen im Osten des heutigen Deutschlands waren ab dem Mittelalter Germanen, die sich gegen die Christianisierung entschieden.
Die Studentenverbindung Skull and Bones beschäftige sich nicht nur mit Okkultismus und Satanismus, sondern habe auch Kontakte zur CIA.
Die Sisson-Dokumente belegen, dass das deutsche Kaiserreich russische Revolutionäre finanziell unterstützte.
Der Mord an dem US-Senator Robert F. Kennedy weist einige Ungereimtheiten in Bezug auf den mutmaßlichen Täter Sirhan Sirhan auf.
Das Siegel der Vereinigten Staaten nimmt Bezug auf den Illuminaten Orden.
Zwischen 1960 und 1980 haben Insassen eines schwarzen Autos der Marke Wolga Kinder verschleppt und gequält, nachdem diese nach der Uhrzeit gefragt wurden.
In Roswell, New Mexico, ist ein UFO abgestürzt.
Reptilohide sind menschenähnliche intelligente Wesen, die von Reptilien oder reptilienartigen Außerirdischen abstammen. Sie haben die Erde unterwandert und sind Teil einer geheimen pyramidenartigen Organisations Struktur.
Das Deutsche Reich besteht juristisch bis heute fort, die Bundesrepublik Deutschland ist dagegen illegal. Eine kommissarische Reichsregierung übt Jurisdiktion über das Deutsche Reich in den Grenzen von 1937 aus.
Das Dritte Reich ist im Besitz eines futuristischen Fluggerätes, der Reichsflugscheibe.
Der Putsch gegen die iranische Regierung Mossadegh galt lange Zeit als Verschwörungstheorie, bis sie im Jahr 2013 von der US-Regierung eingestanden wurde, Codename Operation Ajax.
Das Philadelphia-Experiment des US-Militärs führte dazu, dass ein Kriegsschiff verschwunden ist und an einen anderen Ort teleportiert wurde.
Die US-Regierung oder das US-Militär wussten von dem bevorstehenden Angriff auf Pearl Harbor, unternahmen jedoch nichts, um in der Öffentlichkeit einen Kriegsgrund anführen zu können.
Die Terroranschläge am 11. September 2001 sind von US-amerikanischen Geheimdiensten entweder wissentlich zugelassen oder selbst ausgeführt worden.
Eine Übung unter dem Namen Project Mascal behandelte im Jahr 2000 das Unglück, bei dem ein Flugzeug in das Pentagon stürzt. Einer der Mitwirkenden hat bei American Airlines gearbeitet.
Der Politiker Jürgen Möllemann starb bei einem Fallschirmsprung. Unklar ist, ob es Suizid oder ein Unfall war. Vielleicht war es auch eine gezielte Ermordung?
Die baskische Terrororganisation ETA steckt hinter den Zug Anschlägen in Madrid.
Offiziell heißt es, die Gülen-Bewegung und der Tiefe Staat haben den Putschversuch in der Türkei 2016 durchgeführt. Eine andere Theorie besagt, dass der Putschversuch vom türkischen Staatspräsidenten Recep Tayyip Erdoğan inszeniert worden sei.
Der Amoklauf an der Sandy Hook Elementary School hat nie stattgefunden und wurde stattdessen inszeniert, um die Waffengesetze zu verschärfen.
Unter dem Schlagwort Pizza Gate wurde im Jahr 2016 eine Meldung auf four chan und Reddit verbreitet, die im amerikanischen Präsidentschaftswahlkampf 2016 die Behauptung streute, im Keller einer Pizzeria in Washington, D.C. agiere ein Kinderpornoring, in den auch die Präsidentschaftskandidatin Hillary Clinton verwickelt ist.
Verschwörungstheoretiker wie Alex Jones behaupten, dass es bei dem Attentat in Las Vegas auf das Country-Musik-Festival mehrere Schützen gab.
Es besteht kein kausaler Zusammenhang von H I Viren und der Krankheit AIDS.
HIV ist von den USA entwickelt worden, wurde danach ausgesetzt oder ist entkommen.
Die Einrichtung Area 51 in Nevada beschäftigt und kommuniziert mit außerirdischen Lebensformen.
Das Attentat auf Martin Luther King ist von der US-amerikanischen Regierung ausgeführt oder geplant worden.
Die Teilnehmer der Bilderberg-Konferenz planen eine Weltdiktatur und sind Drahtzieher geschichtsträchtiger Ereignisse, etwa des Irakkriegs.
Kondensstreifen von Flugzeugen enthalten Chemikalien, die sich auf die Bevölkerung auswirken.
Der Vatikan ist im Besitz einer Zeitmaschine.
Das US-amerikanische Forschungsprogramm HAARP ist für Gedankenmanipulation und zur künstlichen Herbeiführung von Naturkatastrophen eingesetzt worden.
Der russische Berg Jamantau beherbergt einen gigantischen unterirdischen Komplex, der militärisch genutzt wird.
Größere Zeiträume der mittelalterlichen Geschichte sind erfunden.
Johannes Paul der Erste ist vergiftet worden, da er interne Machenschaften des Vatikans aufdecken wollte.
Das US-Geheimkomitee Majestic Twelve beschäftigt sich mit dem Wirken von UFOs und Außerirdischen. Und mit Jay Zee Denton.
Geheime, schwarz gekleidete US-Regierungsmitarbeiter sorgen dafür, dass mysteriöse Sichtungen keine Zeugenaussagen hervorbringen.
Die Mondlandungen zwischen 1969 und 1972 durch die NASA haben niemals stattgefunden und wurden lediglich vorgetäuscht.
Marilyn Monroe ist durch einen US-amerikanischen Geheimdienst getötet worden, da sie eine Affäre mit John F. Kennedy hatte.
Im Montauk-Projekt hat das US-Militär versucht, zwischen 1970 und 1990 die Gedanken der Zivilisten zu steuern oder zu beeinflussen.
Paul McCartney von The Beatles ist 1966 gestorben und wurde danach durch einen Doppelgänger ersetzt.
Die Erde ist eine Scheibe. Alle Fotos der NASA sind gefälscht.
Jar Jar Binks ist ein hervorragend ausgebildeter Sith Lord, der seine Trotteligkeit nur vortäuscht.`

var lines = strings.Split(Facts, "\n")

func RandomFact() string {
	randomFact := lines[rand.Intn(len(lines))]
	return strings.TrimSpace(randomFact)
}
