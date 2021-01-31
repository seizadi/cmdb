import { combineReducers } from "redux";

import applicationReducer from "./applicationReducers";
import lifecycleReducer from "./lifecycleReducers";
import environmentReducer from "./environmentReducers";

export default combineReducers( {
  applications: applicationReducer,
  lifecycles: lifecycleReducer,
  environments: environmentReducer,
});
