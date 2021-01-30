import React from "react";
import { BrowserRouter as Router, Route, Switch, Redirect } from "react-router-dom";

// core components
import Admin from "./layouts/Admin.js";

import "./assets/css/material-dashboard-react.css?v=1.9.0";
import ApplicationList from './cmdb/applications/ApplicationList'


function App() {
  return (
    <div className="App">
      <Router>
        <Switch>
          <Route path={'/applications'} exact component={ApplicationList} />
          <Route path="/admin" component={Admin} />
          <Redirect from="/" to="/admin/dashboard" />
        </Switch>
      </Router>
    </div>
  );
}

export default App;
