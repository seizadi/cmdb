import { LIST_CHART_VERSIONS } from "../actions/types";

const INITIAL_STATE = {}

function chartVersionReducer (state = INITIAL_STATE, action) {
  switch(action.type) {
    case LIST_CHART_VERSIONS:
      let newState = {};
      if (action.payload && action.payload.length > 0) {
        action.payload.forEach((chartVersion) => {
          newState = {...newState, [chartVersion.id]: chartVersion}
        });
      }
      return newState;
    default:
      return state;
  }
}

export default chartVersionReducer;
