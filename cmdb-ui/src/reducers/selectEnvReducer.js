import { SELECT_ENVIRONMENT } from "../actions/types";

const INITIAL_STATE = ""

function selectEnvReducer (state = INITIAL_STATE, action) {
  switch(action.type) {
    case SELECT_ENVIRONMENT:
      return action.payload;
    default:
      return state;
  }
}

export default selectEnvReducer;
