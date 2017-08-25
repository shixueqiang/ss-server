import React, {Component} from 'react';

class ProfileEdit extends Component {
    render() {
        return(<div>ProfileEdit {this.props.location.state.Host}</div>);
    };
}
export default ProfileEdit;