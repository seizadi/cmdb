import { combineReducers } from "redux";

import applicationReducer from "./applicationReducers";
import applicationInstanceReducer from "./applicationInstanceReducers";
import lifecycleReducer from "./lifecycleReducers";
import environmentReducer from "./environmentReducers";

export default combineReducers( {
  applications: applicationReducer,
  applicationInstances: applicationInstanceReducer,
  lifecycles: lifecycleReducer,
  environments: environmentReducer,
});
