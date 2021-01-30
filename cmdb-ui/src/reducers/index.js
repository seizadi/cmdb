import { combineReducers } from "redux";

import applicationReducer from "./applicationReducers";

export default combineReducers( {
  applications: applicationReducer,
});
