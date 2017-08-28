import React, {Component} from 'react';
import TextField from 'material-ui/TextField';
import RaisedButton from 'material-ui/RaisedButton';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import $ from 'jquery';
import Dialog from 'material-ui/Dialog';
import FlatButton from 'material-ui/FlatButton';

const style = {
    marginLeft: 60,
    marginTop: 12,
    marginBottom: 12,
};

class ProfileEdit extends Component {
    state = {
        open: false,
        result: "",
        Name: this.props.location.state.Name,
        Host: this.props.location.state.Host,
        LocalPort: this.props.location.state.LocalPort,
        RemotePort: this.props.location.state.RemotePort,
        Password: this.props.location.state.Password,
        Protocol: this.props.location.state.Protocol,
        ProtocolParam: this.props.location.state.ProtocolParam,
        Obfs: this.props.location.state.Obfs,
        ObfsParam: this.props.location.state.ObfsParam,
        Method: this.props.location.state.Method,
        Route: this.props.location.state.Route,
        RemoteDNS: this.props.location.state.RemoteDNS,
        VpnType: this.props.location.state.VpnType,
        Ikev2Type: this.props.location.state.Ikev2Type,
    };

    handleOpen = () => {
        this.setState({open: true,});
    };
    
    handleClose = () => {
        this.setState({open: false,});
        this.props.history.push({pathname:'/toProfiles'});
    };

    handleClick = () => {
        const _this = this;
        $.ajax({
            url:"/profileUpdate",
            data:$("#myform").serialize(),
            type:"POST",
            success: function(data) {
                console.log(data);
                _this.setState({open:true,result: data,});
            }
        });
    }

    handleChange = (event) => {
        var src = event.srcElement || event.target;
        var str = src.name;
        this.setState({
            [str]: src.value,
        });
    };

    render() {
        const actions = [
            <FlatButton
              label="OK"
              primary={true}
              onClick={this.handleClose}
            />,
          ];
        return(
        <div>
            <MuiThemeProvider>
                <div>
                    <form id="myform" action="/profileUpdate" method="post">
                        <input type="hidden" value={this.props.location.state.ID} name="ID"/>
                        <TextField className="profileTextField" floatingLabelText="Name" value={this.state.Name} name="Name" onChange={this.handleChange}/><br />
                        <TextField className="profileTextField" floatingLabelText="Host" value={this.state.Host} name="Host" onChange={this.handleChange}/><br />
                        <TextField className="profileTextField" floatingLabelText="LocalPort" value={this.state.LocalPort} defaultValue="1080" name="LocalPort" onChange={this.handleChange}/><br />
                        <TextField className="profileTextField" floatingLabelText="RemotePort" value={this.state.RemotePort} defaultValue="8388" name="RemotePort" onChange={this.handleChange}/><br />
                        <TextField className="profileTextField" floatingLabelText="Password" value={this.state.Password} type="password" name="Password" onChange={this.handleChange}/><br />
                        <TextField className="profileTextField" floatingLabelText="Protocol" value={this.state.Protocol} defaultValue="origin" name="Protocol" onChange={this.handleChange}/><br />
                        <TextField className="profileTextField" floatingLabelText="ProtocolParam" value={this.state.ProtocolParam} name="ProtocolParam" onChange={this.handleChange}/><br />
                        <TextField className="profileTextField" floatingLabelText="Obfs" value={this.state.Obfs} defaultValue="plain" name="Obfs" onChange={this.handleChange}/><br />
                        <TextField className="profileTextField" floatingLabelText="ObfsParam" value={this.state.ObfsParam} name="ObfsParam" onChange={this.handleChange}/><br />
                        <TextField className="profileTextField" floatingLabelText="Method" value={this.state.Method} defaultValue="aes-256-cfb" name="Method" onChange={this.handleChange}/><br />
                        <TextField className="profileTextField" floatingLabelText="Route" value={this.state.Route} defaultValue="all" name="Route" onChange={this.handleChange}/><br />
                        <TextField className="profileTextField" floatingLabelText="RemoteDNS" value={this.state.RemoteDNS} defaultValue="8.8.8.8" name="RemoteDNS" onChange={this.handleChange}/><br />
                        <TextField className="profileTextField" floatingLabelText="VpnType" value={this.state.VpnType} defaultValue="1" name="VpnType" onChange={this.handleChange}/><br />
                        <TextField className="profileTextField" floatingLabelText="Ikev2Type" value={this.state.Ikev2Type} defaultValue="1" name="Ikev2Type" onChange={this.handleChange}/><br />
                        <RaisedButton label="submit" primary={true} style={style}  onClick={this.handleClick}/>
                    </form>
                    <Dialog
                        actions={actions}
                        modal={false}
                        open={this.state.open}
                        onRequestClose={this.handleClose}
                        >
                        {this.state.result}
                    </Dialog>
                </div>
            </MuiThemeProvider>
        </div>
        );
    };
}
export default ProfileEdit;