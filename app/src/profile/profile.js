import React, {Component} from 'react';
import ReactDOM from 'react-dom';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import AppBar from 'material-ui/AppBar';
import IconButton from 'material-ui/IconButton';
import IconMenu from 'material-ui/IconMenu';
import MenuItem from 'material-ui/MenuItem';
import FlatButton from 'material-ui/FlatButton';
import Toggle from 'material-ui/Toggle';
import MoreVertIcon from 'material-ui/svg-icons/navigation/more-vert';
import NavigationClose from 'material-ui/svg-icons/navigation/close';
import $ from 'jquery';
import './profile.css';
import {
  Table,
  TableBody,
  TableHeader,
  TableHeaderColumn,
  TableRow,
  TableRowColumn,
} from 'material-ui/Table';

export default class TableExampleControlled extends Component {
  state = {
    selected: [],
    data: []
  };

  isSelected = (index) => {
    console.log("isSelected index:" + index);
    return this.state.selected.indexOf(index) !== -1;
  };

  handleRowSelection = (selectedRows) => {
    console.log("selectedRows:" + selectedRows);
    this.setState({
      selected: selectedRows,
    });
  };

  handleResult = (result) => {
    this.setState({
      data: result,
    });
  };

  componentDidMount() {
    const _this = this;
    $.getJSON( "/getAllprofileNotCrypto")
    .done(function( json ) {
      console.log( "JSON Data: " + json.profiles[0].Host);
      _this.handleResult(json.profiles);
    })
    .fail(function( jqxhr, textStatus, error ) {
      var err = textStatus + ", " + error;
      console.log( "Request Failed: " + err );
    });
  }

  render() {
    return (
      <Table onRowSelection={this.handleRowSelection} multiSelectable={true}>
        <TableHeader>
          <TableRow>
            <TableHeaderColumn>ID</TableHeaderColumn>
            <TableHeaderColumn>Host</TableHeaderColumn>
            <TableHeaderColumn>Method</TableHeaderColumn>
            <TableHeaderColumn>Protocol</TableHeaderColumn>
            <TableHeaderColumn>Obfs</TableHeaderColumn>
          </TableRow>
        </TableHeader>
        <TableBody>
          {this.state.data.map((item, index) => 
          <TableRow selected={this.isSelected(index)}>
            <TableRowColumn>{item.ID}</TableRowColumn>
            <TableRowColumn>{item.Host}</TableRowColumn>
            <TableRowColumn>{item.Method}</TableRowColumn>
            <TableRowColumn>{item.Protocol}</TableRowColumn>
            <TableRowColumn>{item.Obfs}</TableRowColumn>
          </TableRow>)}
        </TableBody>
      </Table>
    );
  }
}

class Login extends Component {
    static muiName = 'FlatButton';
  
    render() {
      return (
        <FlatButton {...this.props} label="Login" />
      );
    }
}
  
const Logged = (props) => (
  <IconMenu
    {...props}
    iconButtonElement={
      <IconButton><MoreVertIcon /></IconButton>
    }
    targetOrigin={{horizontal: 'right', vertical: 'top'}}
    anchorOrigin={{horizontal: 'right', vertical: 'top'}}
  >
    <MenuItem primaryText="Edit" />
    <MenuItem primaryText="Remove" />
    <MenuItem primaryText="Sign out" />
  </IconMenu>
);
  
Logged.muiName = 'IconMenu';
  
  /**
   * This example is taking advantage of the composability of the `AppBar`
   * to render different components depending on the application state.
   */
class AppBarExampleComposition extends Component {
  state = {
    logged: true,
  };

  handleChange = (event, logged) => {
    this.setState({logged: logged});
  };

  render() {
    return (
      <div>
        <AppBar
          title="Title"
          iconElementLeft={<IconButton><NavigationClose /></IconButton>}
          iconElementRight={this.state.logged ? <Logged /> : <Login />}
          className="nav"
        />
        <TableExampleControlled/>
      </div>
    );
  }
}

ReactDOM.render(<MuiThemeProvider><AppBarExampleComposition /></MuiThemeProvider>, document.getElementById('root'));