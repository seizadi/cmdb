import { CLEAR_MANIFEST, CREATE_MANIFEST, CLEAR_VALUES, CREATE_VALUES, } from "../actions/types";
import {CircularProgress} from "@material-ui/core";
import React from "react";

const INITIAL_STATE = {artifact: "", values: "", status: <div></div>};

function manifestReducer (state = INITIAL_STATE, action) {
  let newState = {artifact: "No artifact found."};
  switch(action.type) {
    case CLEAR_MANIFEST:
      return {artifact: action.payload, status: <div>Loading <CircularProgress color="secondary"/></div>};
    case CREATE_MANIFEST:
      if (action.payload.artifact) {
        let artifact = action.payload.artifact.replace(/ /g, "\u00a0").split ('\n').map ((item, i) => <div key={i}>{item}</div>);
        newState = {artifact: artifact};
      }
      newState = {...newState, status: <div></div>};
      return newState;
    case CLEAR_VALUES:
      return {values: action.payload, status: <div>Loading <CircularProgress color="secondary"/></div>};
    case CREATE_VALUES:
      if (action.payload.config) {
        let values = action.payload.config.replace(/ /g, "\u00a0").split ('\n').map ((item, i) => <div key={i}>{item}</div>);
        newState = {values: values};
      }
      newState = {...newState, status: <div></div>};
      return newState;
    default:
      return state;
  }
}

export default manifestReducer;
