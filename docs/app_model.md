```mermaid
%%{init: {'theme': 'base', 'themeVariables': { 'primaryColor': '#fffcbb', 'fontFamily': 'aerial', 'fontSize': '50px', 'lineColor': '#ff0000', 'primaryBorderColor': '#ff0000'}}}%%
erDiagram
    APP }o--|| OWNER : "has"
    APP ||--|| DOC : "publish"
    APP ||--|| NAMESPACE : "has"
    APP ||--|| API : "publish"
    APP ||--o{ API : "consume"
    APP ||--|| REPO : "use"
    APP }o--o{ SYNTHETIC : "monitor"       
```
