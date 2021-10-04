
import React from 'react';

export default class PartyForm extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      name: '',
      identifier: ''
    }

    this.handleInputChange = this.handleInputChange.bind(this)
  }

  handleInputChange(event) {
    const target = event.target
    const value = target.type === 'checkbox' ? target.checked : target.value
    const name = target.name

    this.setState({
      [name]: value
    })
  }

  componentDidUpdate() {
    if (this.state.name === '') return
    if (this.state.identifier === '') return

    this.props.onValid(this.state)
  }

  render() {
    const { name } = this.props

    return (
      <form>
        <label>
          {name}: <br/>
          <input
            type="text"
            className="mv2 w-100 pa1"
            placeholder="Name (e.g. The Big Bank)"
            name="name"
            value={this.state.name}
            onChange={this.handleInputChange} />
          <br />
          <input
            type="text"
            className="mv2 w-100 pa1"
            placeholder="Identifier (e.g. TBB)"
            name="identifier"
            value={this.state.identifier}
            onChange={this.handleInputChange} />
        </label>
      </form>
    );
  }
}