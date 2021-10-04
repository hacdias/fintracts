
import React from 'react';
import PartyForm from './PartyForm'

export default class PartiesForm extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
      partyA: null,
      partyB: null
    }
  }

  handlePartyUpdate (name) {
    return (party) => {
      this.setState({
        [name]: party
      })
    }
  }

  render() {
    return (
      <div>
        <PartyForm name="Party A" onValid={this.handlePartyUpdate('partyA')} />
        <PartyForm name="Party B" onValid={this.handlePartyUpdate('partyB')} />
      </div>
    );
  }
}