import { LIST_LIFECYCLES } from "../actions/types";

const INITIAL_STATE = {}

function lifecycleReducer (state = INITIAL_STATE, action) {
  switch(action.type) {
    case LIST_LIFECYCLES:
      let newState = {};
      if (action.payload && action.payload.length > 0) {
        action.payload.forEach((lifecycle) => {
          newState = {...newState, [lifecycle.id]: lifecycle}
        });
      }
      return newState;
    default:
      return state;
  }
}

export default lifecycleReducer;
