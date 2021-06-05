import {SHOW_DISABLED} from "../actions/types";

const INITIAL_STATE = {showDisabled: false}

function graphReducer (state = INITIAL_STATE, action) {
  switch(action.type) {
    case SHOW_DISABLED:
      return {...state, showDisabled: action.payload};
    default:
      return state;
  }
}

export default graphReducer;
