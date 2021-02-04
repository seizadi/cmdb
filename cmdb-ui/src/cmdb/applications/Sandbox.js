/*eslint require-jsdoc: 0, valid-jsdoc: 0, no-console: 0*/
import React from "react";
import defaultConfig from "./graph.config";
import { Graph } from 'react-d3-graph';
import { loadDataset, setValue, merge } from "../../utils/utils";

import "./styles.css";

const sandboxData = loadDataset();

/**
 * This is a sample integration of react-d3-graph, in this particular case all the rd3g config properties
 * will be exposed in a form in order to allow on the fly graph configuration.
 * The data and configuration that are initially loaded can be manipulated via queryParameter on this same
 * Sandbox. You can dynamically load different datasets that are under the `data` folder. If you want
 * for instance to load the data and config under the `small` folder you just need to append "?data=small"
 * to the url when accessing the sandbox.
 */
export default class Sandbox extends React.Component {
  constructor(props) {
    super(props);

    const { config: configOverride, data, fullscreen } = sandboxData;
    const config = Object.assign(defaultConfig, configOverride);
    // TODO: refactor this labelPosition assignment, move to somewhere
    if (config.node.labelPosition === null) {
      config.node.labelPosition = "";
    }

    this.state = {
      config,
      generatedConfig: {},
      data,
      fullscreen,
      nodeIdToBeRemoved: null,
    };
  }

  onClickGraph = () => console.info("Clicked the graph");

  onClickNode = (id, node) => {
    console.info(`Clicked node ${id} in position (${node.x}, ${node.y})`);
    // NOTE: below sample implementation for focusAnimation when clicking on node
    // this.setState({
    //     data: {
    //         ...this.state.data,
    //         focusedNodeId: this.state.data.focusedNodeId !== id ? id : null
    //     }
    // });
  };

  onDoubleClickNode = (id, node) => {
    console.info(`Double clicked node ${id} in position (${node.x}, ${node.y})`);
  };

  onRightClickNode = (event, id, node) => {
    event.preventDefault();
    console.info(`Right clicked node ${id} in position (${node.x}, ${node.y})`);
  };

  onClickLink = (source, target) => console.info(`Clicked link between ${source} and ${target}`);

  onRightClickLink = (event, source, target) => {
    event.preventDefault();
    console.info(`Right clicked link between ${source} and ${target}`);
  };

  onMouseOverNode = (id, node) => {
    console.info(`Do something when mouse is over node ${id} in position (${node.x}, ${node.y})`);
  };

  onMouseOutNode = (id, node) => {
    console.info(`Do something when mouse is out of node ${id} in position (${node.x}, ${node.y})`);
  };

  onMouseOverLink = (source, target) =>
    console.info(`Do something when mouse is over link between ${source} and ${target}`);

  onMouseOutLink = (source, target) =>
    console.info(`Do something when mouse is out of link between ${source} and ${target}`);

  onNodePositionChange = (nodeId, x, y) =>
    console.info(`Node ${nodeId} is moved to new position. New position is (${x}, ${y}) (x,y)`);

  /**
   * Sets on/off fullscreen visualization mode.
   */
  onToggleFullScreen = () => {
    const fullscreen = !this.state.fullscreen;

    this.setState({ fullscreen });
  };

  /**
   * Called when the graph's zoom changes
   * @param {number} prevZoom Previous zoom level
   * @param {number} newZoom New zoom level
   */
  onZoomChange = (prevZoom, newZoom) => {
    this.setState({ currentZoom: newZoom });
  };

  /**
   * Append a new node with some randomness.
   */
  onClickAddNode = () => {
    if (this.state.data.nodes && this.state.data.nodes.length) {
      const maxIndex = this.state.data.nodes.length - 1;
      const minIndex = 0;

      let i = Math.floor(Math.random() * (maxIndex - minIndex + 1) + minIndex),
        nLinks = Math.floor(Math.random() * (5 - minIndex + 1) + minIndex);
      const newNode = `Node ${this.state.data.nodes.length}`;

      this.state.data.nodes.push({ id: newNode });

      while (this.state.data.nodes[i] && this.state.data.nodes[i].id && nLinks) {
        this.state.data.links.push({
          source: newNode,
          target: this.state.data.nodes[i].id,
        });

        i++;
        nLinks--;
      }

      this.setState({
        data: this.state.data,
      });
    } else {
      // 1st node
      const data = {
        nodes: [{ id: "Node 1" }],
        links: [],
      };

      this.setState({ data });
    }
  };

  /**
   * Remove a node.
   */
  onClickRemoveNode = () => {
    if (this.state.data.nodes && this.state.data.nodes.length > 1) {
      const id = this.state.data.nodes[0].id;

      this.state.data.nodes.splice(0, 1);
      const links = this.state.data.links.filter(l => l.source !== id && l.target !== id);
      const data = { nodes: this.state.data.nodes, links };

      this.setState({ data });
    } else {
      console.info("Need to have at least one node!");
    }
  };

  _buildGraphConfig = data => {
    let config = {},
      schemaPropsValues = {};

    for (let k of Object.keys(data.formData)) {
      // Set value mapping correctly for config object of react-d3-graph
      setValue(config, k, data.formData[k]);
      // Set new values for schema of jsonform
      schemaPropsValues[k] = {};
      schemaPropsValues[k]["default"] = data.formData[k];
    }

    return { config, schemaPropsValues };
  };

  refreshGraph = data => {
    const { config, schemaPropsValues } = this._buildGraphConfig(data);

    this.state.schema.properties = merge(this.state.schema.properties, schemaPropsValues);

    this.setState({
      config,
    });
  };

  /**
   * Generate graph configuration file ready to use!
   */
  onSubmit = data => {
    const { config } = this._buildGraphConfig(data);

    this.setState({ generatedConfig: config });
  };

  onClickSubmit = () => {
    // Hack for allow submit button to live outside jsonform
    document.body.querySelector(".invisible-button").click();
  };

  resetGraphConfig = () => {
    const generatedConfig = {};

    this.setState({
      config: defaultConfig,
      generatedConfig,
    });
  };

  /**
   * This function decorates nodes and links with positions. The motivation
   * for this function its to set `config.staticGraph` to true on the first render
   * call, and to get nodes and links statically set to their initial positions.
   * @param  {Object} nodes nodes and links with minimalist structure.
   * @return {Object} the graph where now nodes containing (x,y) coords.
   */
  decorateGraphNodesWithInitialPositioning = nodes => {
    return nodes.map(n =>
      Object.assign({}, n, {
        x: n.x || Math.floor(Math.random() * 500),
        y: n.y || Math.floor(Math.random() * 500),
      })
    );
  };

  /**
   * Before removing elements (nodes, links)
   * from the graph data, this function is executed.
   * https://github.com/oxyno-zeta/react-editable-json-tree#beforeremoveaction
   */
  onBeforeRemoveGraphData = (key, keyPath, deep, oldValue) => {
    if (keyPath && keyPath[0] && keyPath[0] === "nodes" && oldValue && oldValue.id) {
      this.setState({
        nodeIdToBeRemoved: oldValue.id,
      });
    }

    return Promise.resolve();
  };

  /**
   * Update graph data each time an update is triggered
   * by JsonTree
   * @param {Object} data update graph data (nodes and links)
   */
  onGraphDataUpdate = data => {
    const removedNodeIndex = data.nodes.findIndex(n => !n);

    let removedNodeId = null;

    if (removedNodeIndex !== -1 && this.state.nodeIdToBeRemoved) {
      removedNodeId = this.state.nodeIdToBeRemoved;
    }

    const nodes = data.nodes.filter(Boolean);
    const isValidLink = link => link && link.source !== removedNodeId && link.target !== removedNodeId;
    const links = data.links.filter(isValidLink);

    this.setState({
      data: {
        links,
        nodes,
      },
    });
  };

  render() {
    // This does not happens in this sandbox scenario running time, but if we set staticGraph config
    // to true in the constructor we will provide nodes with initial positions
    const data = {
      nodes: this.decorateGraphNodesWithInitialPositioning(this.state.data.nodes),
      links: this.state.data.links,
      focusedNodeId: this.state.data.focusedNodeId,
    };

    const graphProps = {
      id: "graph",
      data,
      config: this.state.config,
      onClickNode: this.onClickNode,
      onDoubleClickNode: this.onDoubleClickNode,
      onRightClickNode: this.onRightClickNode,
      onClickGraph: this.onClickGraph,
      onClickLink: this.onClickLink,
      onRightClickLink: this.onRightClickLink,
      onMouseOverNode: this.onMouseOverNode,
      onMouseOutNode: this.onMouseOutNode,
      onMouseOverLink: this.onMouseOverLink,
      onMouseOutLink: this.onMouseOutLink,
      onNodePositionChange: this.onNodePositionChange,
      onZoomChange: this.onZoomChange,
    };

    if (this.state.fullscreen) {
      graphProps.config = Object.assign({}, graphProps.config, {
        height: window.innerHeight,
        width: window.innerWidth,
      });

      return (
        <div>
          <Graph ref="graph" {...graphProps} />
        </div>
      );
    } else {
      // @TODO: Only show configs that differ from default ones in "Your config" box
      return (
        <div className="container">
          <div className="container__graph">
            <div className="container__graph-area">
              <Graph ref="graph" {...graphProps} />
            </div>
          </div>
        </div>
      );
    }
  }
}
