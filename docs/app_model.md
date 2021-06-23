```mermaid
%%{init: {'theme': 'base', 'themeVariables': { 'primaryColor': '#fffcbb', 'fontFamily': 'verdana', 'fontSize': '30px'}}}%%
erDiagram
    App }o--|| Owner : "has"
    App ||--|| Doc : "publish"
    App ||--|| Namespace : "has"
    App ||--|| Api : "publish"
    App ||--o{ Api : "consume"
    App ||--|| Repo : "use"
    App }o--o{ Synthetic : "monitor"       
```
