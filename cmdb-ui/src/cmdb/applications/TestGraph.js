import React, { useState } from "react";
import { Graph } from 'react-d3-graph';

export default function TestGraph(props) {


  const data = {
    // directed: true,
    nodes: [
      {id: 'chart.v1', color: 'gray', y: 100},
      {id: 'chart.v2', color: 'gray', y: 100},
      {id: 'chart.v3', color: 'gray', y:100},
      {id: 'chart.v4', color: 'gray', y:100},

      {id: 'env-1/sample', y:200},
      {id: 'env-2/sample', y:200},
      {id: 'int-1/sample', y:200},
      {id: 'int-2/sample', y:200},
      {id: 'prod-1/sample', y:200},
      {id: 'prod-2/sample', y:200},

      {id: 'dev', color: 'red', y:300},
      {id: 'integration', color: 'red', y: 300},
      {id: 'prod', color: 'red', y: 300},

    ],
    links: [
      {source: 'dev',
        target: 'integration'
      },
      {source: 'integration', target: 'prod'},
      {source: 'env-1/sample', target: 'dev'},
      {source: 'chart.v4', target: 'env-1/sample'},
      {source: 'env-2/sample', target: 'dev'},
      {source: 'chart.v3', target: 'env-2/sample'},
      {source: 'chart.v3', target: 'int-1/sample'},
      {source: 'int-1/sample', target: 'integration'},
      {source: 'chart.v2', target: 'int-2/sample'},
      {source: 'int-2/sample', target: 'integration'},
      {source: 'chart.v2', target: 'prod-1/sample'},
      {source: 'prod-1/sample', target: 'prod'},
      {source: 'chart.v1', target: 'prod-2/sample'},
      {source: 'prod-2/sample', target: 'prod'},

    ]
  };

// The graph configuration, you only need to pass down properties
// that you want to override, otherwise default ones will be used
  const myConfig = {
    // height: 400,
    // width: 800,
    nodeHighlightBehavior: true,
    disableLinkForceConfig: false, // --> new API config
    node: {
      color: 'lightgreen',
      size: 120,
      highlightStrokeColor: 'blue'
    },
    link: {
      highlightColor: 'lightblue'
    }
  };

//   const myConfig = {
//     // 'automaticRearrangeAfterDropNode': false,
//     'collapsible': false,
//     'directed': true,
//     'focusAnimationDuration': 0.75,
//     'height': 700,
//     'highlightDegree': 2,
//     'highlightOpacity': 0.2,
//     'linkHighlightBehavior': true,
//     'maxZoom': 12,
//     'minZoom': 0.05,
//     'nodeHighlightBehavior': true,
//     'panAndZoom': false,
//     'staticGraph': false,
//     'staticGraphWithDragAndDrop': true,
//     'd3': {
//       'alphaTarget': 0.05,
//       'gravity': -100,
//       'linkLength': 150,
//       'linkStrength': 1
//     },
//     'node': {
//       'mouseCursor': 'grab',
//       'opacity': 1,
//       'renderLabel': true,
//       'size': 1500,
//       // 'viewGenerator': node => <QuestionNodeSvg question={node}/>
//     },
//     'link': {
//       'color': '#000',
//       'highlightColor': 'red',
//       'mouseCursor': 'pointer',
//       'opacity': 1,
//       'semanticStrokeWidth': false,
//       'strokeWidth': 2
//     }
//   };

  console.log("data....", data)
// Callback to handle click on the graph.
// @param {Object} event click dom event
  const onClickGraph = function(event) {
    // window.alert('Clicked the graph background');
  };

  const onClickNode = function(nodeId, node) {
    // window.alert('Clicked node ${nodeId} in position (${node.x}, ${node.y})');
  };

  const onDoubleClickNode = function(nodeId, node) {
    // window.alert('Double clicked node ${nodeId} in position (${node.x}, ${node.y})');
  };

  const onRightClickNode = function(event, nodeId, node) {
    // window.alert('Right clicked node ${nodeId} in position (${node.x}, ${node.y})');
  };

  const onMouseOverNode = function(nodeId, node) {
    console.info(`Mouse over node ${nodeId} in position (${node.x}, ${node.y})`);
  };

  const onMouseOutNode = function(nodeId, node) {
    // window.alert(`Mouse out node ${nodeId} in position (${node.x}, ${node.y})`);
  };

  const onClickLink = function(source, target) {
    // window.alert(`Clicked link between ${source} and ${target}`);
  };

  const onRightClickLink = function(event, source, target) {
    // window.alert('Right clicked link between ${source} and ${target}');
  };

  const onMouseOverLink = function(source, target) {
    // console.info(`Mouse over in link between ${source} and ${target}`);
  };

  const onMouseOutLink = function(source, target) {
    // window.alert(`Mouse out link between ${source} and ${target}`);
  };

  const onNodePositionChange = function(nodeId, x, y) {
    // console.log(`Node ${nodeId} moved to new position x= ${x} y= ${y}`);
  };

// Callback that's called whenever the graph is zoomed in/out
// @param {number} previousZoom the previous graph zoom
// @param {number} newZoom the new graph zoom
  const onZoomChange = function(previousZoom, newZoom) {
    // window.alert(`Graph is now zoomed at ${newZoom} from ${previousZoom}`);
  };

  return (
    <>
      <Graph
        id='graph-id' // id is mandatory, if no id is defined rd3g will throw an error
        data={data}
        config={myConfig}
        onClickGraph={onClickGraph}
        onClickNode={onClickNode}
        onDoubleClickNode={onDoubleClickNode}
        onRightClickNode={onRightClickNode}
        onClickLink={onClickLink}
        onRightClickLink={onRightClickLink}
        onMouseOverNode={onMouseOverNode}
        onMouseOutNode={onMouseOutNode}
        onMouseOverLink={onMouseOverLink}
        onMouseOutLink={onMouseOutLink}
        onNodePositionChange={onNodePositionChange}
        onZoomChange={onZoomChange}/>
    </>
  );
}
