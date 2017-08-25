import React, {Component} from 'react';
import TextField from 'material-ui/TextField';
import RaisedButton from 'material-ui/RaisedButton';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';

const style = {
    marginLeft: 60,
    marginTop: 12,
    marginBottom: 12,
};

class ProfileInsert extends Component {
    handleClick = () => {
        console.log("submit");
    }

    render() {
        return(
        <div>
            <MuiThemeProvider>
                <form action="/profileUpdate" method="post">
                    <TextField className="profileTextField" floatingLabelText="Name" name="Name"/><br />
                    <TextField className="profileTextField" floatingLabelText="Host" name="Host"/><br />
                    <TextField className="profileTextField" floatingLabelText="LocalPort" defaultValue="1080" name="LocalPort"/><br />
                    <TextField className="profileTextField" floatingLabelText="RemotePort" defaultValue="8388" name="RemotePort"/><br />
                    <TextField className="profileTextField" floatingLabelText="Password" type="password" name="Password"/><br />
                    <TextField className="profileTextField" floatingLabelText="Protocol" defaultValue="origin" name="Protocol"/><br />
                    <TextField className="profileTextField" floatingLabelText="ProtocolParam" name="ProtocolParam"/><br />
                    <TextField className="profileTextField" floatingLabelText="Obfs" defaultValue="plain" name="Obfs"/><br />
                    <TextField className="profileTextField" floatingLabelText="ObfsParam" name="ObfsParam"/><br />
                    <TextField className="profileTextField" floatingLabelText="Method" defaultValue="aes-256-cfb" name="Method"/><br />
                    <TextField className="profileTextField" floatingLabelText="Route" defaultValue="all" name="Route"/><br />
                    <TextField className="profileTextField" floatingLabelText="RemoteDNS" defaultValue="8.8.8.8" name="RemoteDNS"/><br />
                    <TextField className="profileTextField" floatingLabelText="VpnType" defaultValue="1" name="VpnType"/><br />
                    <TextField className="profileTextField" floatingLabelText="Ikev2Type" defaultValue="1" name="Ikev2Type"/><br />
                    <RaisedButton label="submit" primary={true} style={style}  onClick={this.handleClick}/>
                </form>
            </MuiThemeProvider>
        </div>
        );
    };
}
export default ProfileInsert;