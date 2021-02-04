/*eslint require-jsdoc: 0, valid-jsdoc: 0, no-undef: 0, no-empty: 0, no-console: 0*/
import React from "react";

const LABEL_POSITION_OPTIONS = ["left", "right", "top", "bottom", "center"];

const LINE_TYPES = {
  STRAIGHT: "STRAIGHT",
  CURVE_SMOOTH: "CURVE_SMOOTH",
  CURVE_FULL: "CURVE_FULL",
};

/**
 * This two functions generate the react-jsonschema-form
 * schema from some passed graph configuration.
 */
function formMap(k, v) {
  // customized props
  switch (k) {
    case "link.type": {
      return {
        type: "array",
        title: "link.type",
        items: {
          enum: Object.keys(LINE_TYPES),
        },
        uniqueItems: true,
      };
    }
    case "node.labelPosition": {
      return {
        type: "array",
        title: "node.labelPosition",
        items: {
          enum: LABEL_POSITION_OPTIONS,
        },
        uniqueItems: true,
      };
    }
  }

  return {
    title: k,
    type: typeof v,
    default: v,
  };
}

function generateFormSchema(o, rootSpreadProp, accum = {}) {
  for (let k of Object.keys(o)) {
    const kk = rootSpreadProp ? `${rootSpreadProp}.${k}` : k;

    if (o[k] !== undefined && o[k] !== null && typeof o[k] !== "function") {
      typeof o[k] === "object" ? generateFormSchema(o[k], kk, accum) : (accum[kk] = formMap(kk, o[k]));
    }
  }

  return accum;
}

function loadDataset() {

  const fullscreen = false;

  // TODO - Play around with more exotic dataset later
  // if (queryParams && queryParams.data) {
  //   const dataset = queryParams.data.toLowerCase();
  //
  //   try {
  //     const data = require(`./data/${dataset}/${dataset}.data`);
  //     const datasetConfig = require(`./data/${dataset}/${dataset}.config`);
  //     // hasOwnProperty(datasetConfig, "default") hack to get around mixed module systems
  //     const tmp = Object.prototype.hasOwnProperty.call(datasetConfig, "default")
  //       ? datasetConfig.default
  //       : datasetConfig;
  //     const config = merge(DEFAULT_CONFIG, tmp);
  //
  //     return { data, config, fullscreen };
  //   } catch (error) {
  //     console.warn(
  //       `dataset with name ${dataset} not found, falling back to default, make sure it is a valid dataset`
  //     );
  //   }
  // }

  const config = {};
  const data = require("../cmdb/applications/data/default");

  return {
    config,
    data,
    fullscreen,
  };
}

function isArray(o) {
  return o && typeof o === "object" && Object.prototype.hasOwnProperty.call(o, "length");
}

function setValue(obj, access, value) {
  if (typeof access == "string") {
    access = access.split(".");
  }

  // check for non existence of root property before advancing
  if (!obj[access[0]]) {
    obj[access[0]] = {};
  }

  const v = isArray(value) ? value[0] : value;

  if (access.length > 1) {
    setValue(obj[access.shift()], access, v);
  } else {
    obj[access[0]] = v;
  }
}

/**
 * This function merges two objects o1 and o2, where o2 properties override existent o1 properties, and
 * if o2 doesn't posses some o1 property the fallback will be the o1 property.
 * @param  {Object} o1 - object.
 * @param  {Object} o2 - object that will override o1 properties.
 * @param  {int} [_depth=0] - the depth at which we are merging the object.
 * @returns {Object} object that is the result of merging o1 and o2, being o2 properties priority overriding
 * existent o1 properties.
 * @memberof utils
 */
function merge(o1 = {}, o2 = {}, _depth = 0) {
  let o = {};

  if (Object.keys(o1 || {}).length === 0) {
    return o2 && !isEmptyObject(o2) ? o2 : {};
  }

  for (let k of Object.keys(o1)) {
    const nestedO = !!(o2[k] && typeof o2[k] === "object" && typeof o1[k] === "object" && _depth < MAX_DEPTH);

    if (nestedO) {
      const r = merge(o1[k], o2[k], _depth + 1);

      o[k] =
        Object.prototype.hasOwnProperty.call(o1[k], "length") && Object.prototype.hasOwnProperty.call(o2[k], "length")
          ? Object.keys(r).map(rk => r[rk])
          : r;
    } else {
      o[k] = Object.prototype.hasOwnProperty.call(o2, k) ? o2[k] : o1[k];
    }
  }

  return o;
}

export { generateFormSchema, loadDataset, setValue, merge };
