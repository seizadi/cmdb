import { LIST_ENVIRONMENTS } from "../actions/types";

const INITIAL_STATE = {}

function environmentReducer (state = INITIAL_STATE, action) {
  switch(action.type) {
    case LIST_ENVIRONMENTS:
      let newState = {};
      if (action.payload && action.payload.length > 0) {
        action.payload.forEach((environment) => {
          newState = {...newState, [environment.id]: environment}
        });
      }
      return newState;
    default:
      return state;
  }
}

export default environmentReducer;
