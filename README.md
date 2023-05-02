# Beispielcode/Aufgaben zu verketteten Listen

## Aufgaben zur vorhandenen Liste von Zahlen

Im Package `dllist_int` ist eine Implementierung einer doppelt verketteten Liste
vorgegeben, deren Element Zahlen als Werte enthalten.
Es sind noch einige Implementierungen von Methoden/Funktionen zu vervollständigen:

* Die Funktion `swapNodes` erwartet zwei Knoten aus einer Liste und soll sie vertauschen.
* Die Methoden `PushBack`, `PushFront`, `PopBack` und `PopFront` aus dem Datentyp `DLList`
  sollen die entsprechenden Operationen auf der Liste ausführen.
* Die Methode `Get`aus dem Datentyp `DLList` soll den Wert des Elements an einer
  bestimmten Position zurückgeben.
* Die Methode `Swap`aus dem Datentyp `DLList` soll zwei Knoten aus der Liste vertauschen.

*Hinweis*: Im Package `dllist_int_solution' ist eine Lösung zu finden.

## Erstellung einer Liste für einen komplexeren Datentyp

Erstellen Sie auf Basis des Packages `dllist_int` ein Package für einen Listen-Datentyp,
dessen Elemente nicht nur Zahlen, sondern komplexere Werte enthalten.

**Einfache Beispiele/Ideen für Daten anstelle von Zahlen:**

* Einträge in einem Telefonbuch mit Namen, Telefon, E-Mail, ...
* Artikel mit Artikelnummer, Bezeichnung, Preis, ...
* Kinofilme mit Titel, Erscheinungsjahr, Genre, Bewertungen ...
* Personen in einem sozialen Netzwerk mit Name sowie jeweils einer Liste von Freunden.

*Hinweis:* Achten Sie darauf, Ihren Datentyp zunächst möglichst einfach zu halten.
Auch wenn Sie vielleicht sofort eine Reihe an Ideen haben, sollten Sie zunächst
nur das Minumum umsetzen, um die Funktionalität der Liste zu testen.
Erweitern Sie Ihren Datentyp dann schrittweise um weitere Eigenschaften.
