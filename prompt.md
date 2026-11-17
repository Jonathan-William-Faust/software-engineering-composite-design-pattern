Implementiere das **Composite Design Pattern** in Go gemäß folgender Spezifikation.

### 1. Datenstrukturen

Implementiere folgende Structs:

- `Window`
- `Frame`
- `Glass`

Definiere zusätzlich das Interface `Item`, das von allen drei Structs implementiert wird.

**Zusammensetzungsregeln:**

- Ein `Window` enthält entweder genau ein `Glass` oder genau vier weitere `Window`-Objekte und einen `Frame`.
- `Frame` und `Glass` sind Blattelemente und können keine weiteren Items enthalten.
- Stelle sicher, dass die Zusammensetzungsregeln bei der Erstellung der Objekte eingehalten werden.

### 2. Interface Item

Definiere das Interface mit folgenden Methoden:

- `Price() float64`: Gibt den Gesamtpreis des Items zurück. Bei `Frame` und `Glass` wird ein festgelegter Preis zurückgegeben. Bei `Window` entspricht der Preis der Summe aller enthaltenen Komponenten.
- `Display() string`: Gibt eine hierarchische String-Repräsentation des Items und seiner Bestandteile zurück. Die Ausgabe soll die Baumstruktur abbilden. Preise sollen ausschließlich bei den Blattelementen (`Frame` und `Glass`) angezeigt werden.

### 3. Preise

Verwende folgende Festpreise:

- `Frame`: 50.00
- `Glass`: 30.00

Die Preise eines `Window` ergeben sich rekursiv aus der Summe der Preise seiner direkten und indirekten Komponenten.

### 4. Demonstrationsbeispiel

Erstelle ein kleines ausführbares Beispiel (`main`), das:

- Ein `Window` mit einem `Glass` erstellt.
- Ein weiteres `Window` mit vier untergeordneten `Window`-Objekten und einem `Frame` erstellt:
  - Drei der vier untergeordneten Windows enthalten jeweils genau ein `Glass`.
  - Das vierte untergeordnete Window enthält wiederum vier untergeordnete `Window`-Objekte.
  - Jedes dieser vier Windows der zweiten hierarchischen Stufe enthält genau ein `Glass`.
- Die Methoden `Price()` und `Display()` für beide Window-Objekte aufruft und die Ergebnisse ausgibt.

### 5. Implementierungsvorgaben

- Verwende idiomatisches Go.
- Nutze das Interface `Item` konsequent zur Umsetzung des Composite Patterns.
- Vermeide unnötige Abstraktionen und zusätzliche Dependencies.
- Füge kurze, sachliche Kommentare zur Erläuterung der wesentlichen Implementierungsentscheidungen ein.
- Beschreibe die Methoden `Price()` und `Display()` direkt im Interface durch aussagekräftige Go-Kommentare.
- Achte auf eine klare Trennung zwischen Composite (`Window`) und Leaf-Elementen (`Frame`, `Glass`).
- Stelle sicher, dass der Code vollständig kompilierbar und ausführbar ist.

Gib den vollständigen Go-Quellcode aus und demonstriere die erwarteten Ausgaben.