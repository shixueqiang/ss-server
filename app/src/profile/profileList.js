import React, {Component} from 'react';
import AppBar from 'material-ui/AppBar';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import IconButton from 'material-ui/IconButton';
import IconMenu from 'material-ui/IconMenu';
import MenuItem from 'material-ui/MenuItem';
import FlatButton from 'material-ui/FlatButton';
import MoreVertIcon from 'material-ui/svg-icons/navigation/more-vert';
import NavigationClose from 'material-ui/svg-icons/navigation/close';
import { Table } from 'antd';
import 'antd/dist/antd.css';
import $ from 'jquery';
import './profile.css';

// export default class TableExampleControlled extends Component {
//   state = {
//     selected: [],
//     data: []
//   };

//   isSelected = (index) => {
//     console.log("isSelected index:" + index);
//     return this.state.selected.indexOf(index) !== -1;
//   };

//   handleRowSelection = (selectedRows) => {
//     console.log("selectedRows:" + selectedRows);
//     this.setState({
//       selected: selectedRows,
//     });
//   };

//   handleResult = (result) => {
//     this.setState({
//       data: result,
//     });
//   };

//   componentDidMount() {
//     const _this = this;
//     $.getJSON( "http://127.0.0.1:8055/getAllprofileNotCrypto")
//     .done(function( json ) {
//       console.log( "JSON Data: " + json.profiles[0].Host);
//       _this.handleResult(json.profiles);
//     })
//     .fail(function( jqxhr, textStatus, error ) {
//       var err = textStatus + ", " + error;
//       console.log( "Request Failed: " + err );
//     });
//   }

//   render() {
//     return (
//       <Table onRowSelection={this.handleRowSelection} multiSelectable={true}>
//         <TableHeader>
//           <TableRow>
//             <TableHeaderColumn>ID</TableHeaderColumn>
//             <TableHeaderColumn>Host</TableHeaderColumn>
//             <TableHeaderColumn>Method</TableHeaderColumn>
//             <TableHeaderColumn>Protocol</TableHeaderColumn>
//             <TableHeaderColumn>Obfs</TableHeaderColumn>
//           </TableRow>
//         </TableHeader>
//         <TableBody>
//           {this.state.data.map((item, index) => 
//           <TableRow key={index} selected={this.isSelected(index)}>
//             <TableRowColumn>{item.ID}</TableRowColumn>
//             <TableRowColumn>{item.Host}</TableRowColumn>
//             <TableRowColumn>{item.Method}</TableRowColumn>
//             <TableRowColumn>{item.Protocol}</TableRowColumn>
//             <TableRowColumn>{item.Obfs}</TableRowColumn>
//           </TableRow>)}
//         </TableBody>
//       </Table>
//     );
//   }
// }

class Login extends Component {
    static muiName = 'FlatButton';
  
    render() {
      return (
        <FlatButton {...this.props} label="Login" />
      );
    }
}

class Logged extends Component {

  handleChangeSingle = (event, value) => {
    switch(value) {
      case '1':
        window.location.href="/profileEdit"
      break;
      case '2':
      break;
      case '3':
      break;
      default :

    }
  };

  render() {
    return (
      <IconMenu {...this.props}  
      iconButtonElement={
        <IconButton><MoreVertIcon /></IconButton>
      }
      targetOrigin={{horizontal: 'right', vertical: 'top'}}
      anchorOrigin={{horizontal: 'right', vertical: 'top'}}
      onChange={this.handleChangeSingle}>
      <MenuItem value="1" primaryText="Edit" />
      <MenuItem value="2" primaryText="Remove" />
      <MenuItem value="3" primaryText="Sign out" />
    </IconMenu>
    );
  };
}
Logged.muiName = 'IconMenu';
  
  /**
   * This example is taking advantage of the composability of the `AppBar`
   * to render different components depending on the application state.
   */
class ProfileList extends Component {
  state = {
    logged: true,
    data: [],
  };

  handleChange = (event, logged) => {
    this.setState({logged: logged});
  };

  editProfile = () => {

  };

  updateData = (profiles) => {
      console.log(profiles.length);
  };

  render() {
    return (
      <div>
        <MuiThemeProvider>
            <AppBar
            title="VPN PROFILE"
            iconElementLeft={<IconButton><NavigationClose/></IconButton>}
            iconElementRight={this.state.logged ? <Logged callback={this.editProfile}/> : <Login />}
            className="nav"/>
        </MuiThemeProvider>
        <ExampleTable callback={this.updateData}/>  
      </div>
    );
  }
}

const columns = [{
  title: 'ID',
  dataIndex: 'ID',
}, {
  title: 'Host',
  dataIndex: 'Host',
}, {
  title: 'Method',
  dataIndex: 'Method',
}, {
  title: 'Protocol',
  dataIndex: 'Protocol',
}, {
  title: 'Obfs',
  dataIndex: 'Obfs',
}];

class ExampleTable extends React.Component {
  state = {
    selectedRowKeys: [],  // Check here to configure the default column
    loading: false,
    data: [],
  };

  componentDidMount() {
    const _this = this;
    $.getJSON( "http://127.0.0.1:8055/getAllprofileNotCrypto")
    .done(function( json ) {
      console.log( "JSON Data: " + json.profiles[0].Host);
      _this.setState({data: json.profiles,});
    })
    .fail(function( jqxhr, textStatus, error ) {
      var err = textStatus + ", " + error;
      console.log( "Request Failed: " + err );
    });
  }

  onSelectChange = (selectedRowKeys) => {
    console.log('selectedRowKeys changed: ', selectedRowKeys);
    this.setState({ selectedRowKeys });
    var array = new Array();
    for(var i in selectedRowKeys) {
        array.push(this.state.data[i]);
    }
    this.props.callback(array);
  }
  render() {
    const { loading, selectedRowKeys } = this.state;
    const rowSelection = {
      selectedRowKeys,
      onChange: this.onSelectChange,
    };
    const hasSelected = selectedRowKeys.length > 0;
    return (
      <div>
        <Table rowSelection={rowSelection} columns={columns} dataSource={this.state.data} />
      </div>
    );
  }
}

export default ProfileList;