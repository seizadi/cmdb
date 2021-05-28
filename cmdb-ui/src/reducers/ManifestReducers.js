import { CLEAR_MANIFEST, CREATE_MANIFEST } from "../actions/types";
import {CircularProgress} from "@material-ui/core";
import React from "react";

const INITIAL_STATE = {artifact: "", status: <div></div>};

function manifestReducer (state = INITIAL_STATE, action) {
  let newState = {artifact: "No artifact found."};
  switch(action.type) {
    case CLEAR_MANIFEST:
      return {artifact: action.payload, status: <div>Loading <CircularProgress color="secondary"/></div>};
    case CREATE_MANIFEST:
      if (action.payload.config) {
        let artifact = action.payload.config.replace(/ /g, "\u00a0").split ('\n').map ((item, i) => <div key={i}>{item}</div>);
        newState = {artifact: artifact};
      }
      newState = {...newState, status: <div></div>};
      return newState;
    default:
      return state;
  }
}

export default manifestReducer;
