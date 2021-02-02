import { LIST_APPLICATION_INSTANCES } from "../actions/types";

const INITIAL_STATE = {}

function applicationInstanceReducer (state = INITIAL_STATE, action) {
  switch(action.type) {
    case LIST_APPLICATION_INSTANCES:
      let newState = {};
      if (action.payload && action.payload.length > 0) {
        action.payload.forEach((applicationInstance) => {
          newState = {...newState, [applicationInstance.id]: applicationInstance}
        });
      }
      return newState;
    default:
      return state;
  }
};

export default applicationInstanceReducer;
