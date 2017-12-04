const {connect} = window['react-redux'];
const {bindActionCreators} = window.redux;

import {startMeeting} from '../../actions';

import ChannelHeaderButton from './channel_header_button.jsx';

function mapStateToProps(state, ownProps) {
    return {
        channelId: state.entities.channels.currentChannelId,
        ...ownProps
    };
}

function mapDispatchToProps(dispatch) {
    return {
        actions: bindActionCreators({
            startMeeting
        }, dispatch)
    };
}

export default connect(mapStateToProps, mapDispatchToProps)(ChannelHeaderButton);