import React, {Component} from 'react';
import TextField from 'material-ui/TextField';
import RaisedButton from 'material-ui/RaisedButton';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import './profile.css';

const style = {
    marginLeft: 60,
    marginTop: 12,
    marginBottom: 12,
};

class ProfileEdit extends Component {
    handleClick = () => {
        console.log("submit");
    }

    render() {
        return(
        <div>
            <MuiThemeProvider>
                <form action="/profileUpdate" method="post">
                    <TextField className="profileTextField" floatingLabelText="Name" value={this.props.location.state.Name} name="Name"/><br />
                    <TextField className="profileTextField" floatingLabelText="Host" value={this.props.location.state.Host} name="Host"/><br />
                    <TextField className="profileTextField" floatingLabelText="LocalPort" value={this.props.location.state.LocalPort} defaultValue="1080" name="LocalPort"/><br />
                    <TextField className="profileTextField" floatingLabelText="RemotePort" value={this.props.location.state.RemotePort} defaultValue="8388" name="RemotePort"/><br />
                    <TextField className="profileTextField" floatingLabelText="Password" value={this.props.location.state.Password} type="password" name="Password"/><br />
                    <TextField className="profileTextField" floatingLabelText="Protocol" value={this.props.location.state.Protocol} defaultValue="origin" name="Protocol"/><br />
                    <TextField className="profileTextField" floatingLabelText="ProtocolParam" value={this.props.location.state.ProtocolParam} name="ProtocolParam"/><br />
                    <TextField className="profileTextField" floatingLabelText="Obfs" value={this.props.location.state.Obfs} defaultValue="plain" name="Obfs"/><br />
                    <TextField className="profileTextField" floatingLabelText="ObfsParam" value={this.props.location.state.ObfsParam} name="ObfsParam"/><br />
                    <TextField className="profileTextField" floatingLabelText="Method" value={this.props.location.state.Method} defaultValue="aes-256-cfb" name="Method"/><br />
                    <TextField className="profileTextField" floatingLabelText="Route" value={this.props.location.state.Route} defaultValue="all" name="Route"/><br />
                    <TextField className="profileTextField" floatingLabelText="RemoteDNS" value={this.props.location.state.RemoteDNS} defaultValue="8.8.8.8" name="RemoteDNS"/><br />
                    <TextField className="profileTextField" floatingLabelText="VpnType" value={this.props.location.state.vpnType} defaultValue="1" name="VpnType"/><br />
                    <TextField className="profileTextField" floatingLabelText="Ikev2Type" value={this.props.location.state.Ikev2Type} defaultValue="1" name="Ikev2Type"/><br />
                    <RaisedButton label="submit" primary={true} style={style}  onClick={this.handleClick}/>
                </form>
            </MuiThemeProvider>
        </div>
        );
    };
}
export default ProfileEdit;