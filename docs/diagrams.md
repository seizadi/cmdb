# References
* [How to get diagrams in MkDocs](https://chrieke.medium.com/the-best-mkdocs-plugins-and-customizations-fc820eb19759)
* [How to setup MkDocs on Mac and Github](https://suedbroecker.net/2021/01/25/how-to-install-mkdocs-on-mac-and-setup-the-integration-to-github-pages/)
* [Mermaid](https://mermaid-js.github.io/mermaid)
* [Mermaid2 plugin](https://github.com/fralau/mkdocs-mermaid2-plugin#declaring-the-superfences-extension)
* [Mermaid Live Editor](https://mermaid-js.github.io/mermaid-live-editor/#/edit/eyJjb2RlIjoiZ3JhcGggVERcbiAgICBBW0NocmlzdG1hc10gLS0-fEdldCBtb25leXwgQihHbyBzaG9wcGluZylcbiAgICBCIC0tPiBDe0xldCBtZSB0aGlua31cbiAgICBDIC0tPnxPbmV8IERbTGFwdG9wXVxuICAgIEMgLS0-fFR3b3wgRVtpUGhvbmVdXG4gICAgQyAtLT58VGhyZWV8IEZbZmE6ZmEtY2FyIENhcl0iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOmZhbHNlfQ)
* [Mermaid Diagram Syntax](https://mermaid-js.github.io/mermaid/#/flowchart?id=flowcharts-basic-syntax)

## Class or UML

Class Diagram:

The different cardinality options are:

| Value  | Meaning                        |
| :---:  | ----:                         |
| 1	     | Exactly one                   |
| 0..1	 | Zero or one                   |
| 1..*   | One or more (no upper limit)  |
| 0..n	 | Zero to n {where n>1}         |
| 1..n	 | One to n {where n>1}          |
| *      | Many                          |
| n      | n {where n>1}                 |
  
```mermaid
classDiagram
    CloudProvider "1" --> "*" Region
    Region "1" --> "1..*" Network
    Network "1" --> "*" KubeCluster
    KubeCluster "1" --> "*" Deployment
    Application "1" --> "1..*" AppInstance
    Application "1" --> "1..*" ChartVersion
    AppInstance "1" --> "1" ChartVersion
    AppInstance "1" --> "1" Deployment
```

ER Diagram:

| Value (left) | Value (right) |Meaning                        |
| :---         | :---:         | ----:                         |
| &#124;o      | o&#124;	   | Zero or one                   |
| &#124;&#124;| &#124;&#124;   | Exactly one                   |
| }o	       | o{	           | Zero or more (no upper limit) |
| }&#124;      | &#124;{	   | One or more (no upper limit)  |



```mermaid
erDiagram
    CloudProvider {
        string account
    }
    CloudProvider ||--|{ Region : ""
    Region ||--o{ Network : ""
    Network ||--o{ KubeCluster : ""
    KubeCluster ||--o{ Deployment : ""
    Application ||--o{ AppInstance : ""
    Application ||--o{ ChartVersion : ""
    AppInstance ||--|| ChartVersion : ""
    AppInstance ||--|| Deployment : ""
```

## Graph or Flowchart

```mermaid
graph LR
A((Client)) --> B[Load Balancer]
B --> C[Server01] --> E[Application1] & F[Application2]
B --> D[Server02] --> G[(Database)]
```

```mermaid
graph TB
A((Start)) --o B
subgraph stage one
B[Do Something] --> C{Did it work?}
C -- No --> B
end
C -- Yes --> D[Do more!]
subgraph stage two
D --> E{Are we done?}
E -- No --> B
E -- Yes --> F[Finish up]
end
F ----x G((Stop))
```

## Sequence Diagram

Simple...

```mermaid
sequenceDiagram
    Alice->>John: Hello John, how are you?
    John-->>Alice: Great!
    Alice-)John: See you later!
```

Complex....

```mermaid
%% Example of sequence diagram
  sequenceDiagram
    Alice->>Bob: Hello Bob, how are you?
    alt is sick
    Bob->>Alice: Not so good :(
    else is well
    Bob->>Alice: Feeling fresh like a daisy
    end
    opt Extra response
    Bob->>Alice: Thanks for asking
    end
```

More Complex....
```mermaid
  sequenceDiagram
    autonumber
    rect rgb(0, 255, 0)
        par Alice to Bob
            Alice->>Bob: Hello Bob, how are you?
        and Alice to John
            Alice->>John: Hello John, how are you?
            rect rgba(0, 0, 255, .3)
                par John to Jane
                Note over John: Note we embed another parallel here
                    John->>John: Should I tell Jane?
                    John->>Jane: Alice called me!
                    John->>Alice: I'm doing good
                end
            end
        and Alice to Jane
            Alice->>Jane: Hello Jane, how are you?
        end
    end
```

## Journey

```mermaid
journey
    title My working day
    section Go to work
      Make tea: 5: Me
      Go upstairs: 3: Me
      Do work: 1: Me, Cat
    section Go home
      Go downstairs: 5: Me
      Sit down: 5: Me
```

## Gantt Chart

```mermaid
gantt
    dateFormat  YYYY-MM-DD
    title       Adding GANTT diagram functionality to mermaid
    excludes    weekends
    %% (`excludes` accepts specific dates in YYYY-MM-DD format, days of the week ("sunday") or "weekends", but not the word "weekdays".)

    section A section
    Completed task            :done,    des1, 2014-01-06,2014-01-08
    Active task               :active,  des2, 2014-01-09, 3d
    Future task               :         des3, after des2, 5d
    Future task2              :         des4, after des3, 5d

    section Critical tasks
    Completed task in the critical line :crit, done, 2014-01-06,24h
    Implement parser and jison          :crit, done, after des1, 2d
    Create tests for parser             :crit, active, 3d
    Future task in critical line        :crit, 5d
    Create tests for renderer           :2d
    Add to mermaid                      :1d

    section Documentation
    Describe gantt syntax               :active, a1, after des1, 3d
    Add gantt diagram to demo page      :after a1  , 20h
    Add another diagram to demo page    :doc1, after a1  , 48h

    section Last section
    Describe gantt syntax               :after doc1, 3d
    Add gantt diagram to demo page      :20h
    Add another diagram to demo page    :48h
```

## Pie Chart

```mermaid
pie
    title Key elements in Product X
    "Calcium" : 42.96
    "Potassium" : 50.05
    "Magnesium" : 10.01
    "Iron" :  5
```
