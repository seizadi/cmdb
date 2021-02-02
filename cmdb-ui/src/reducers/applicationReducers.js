import { LIST_APPLICATIONS } from "../actions/types";

const INITIAL_STATE = {}

function applicationReducer (state = INITIAL_STATE, action) {
  switch(action.type) {
    case LIST_APPLICATIONS:
      let newState = {};
      if (action.payload && action.payload.length > 0) {
        action.payload.forEach((application) => {
          newState = {...newState, [application.id]: application}
        });
      }
      return newState;
    default:
      return state;
  }
}

export default applicationReducer;
