import React, { Component } from "react";
import PropTypes from "prop-types";
import {
  Box,
  Heading,
  Select,
  Form,
  Button,
  TextInput,
  Text,
  FormField
} from "grommet";

export default class NewErrorForm extends Component {
  static propTypes = {
    onSave: PropTypes.func
  };

  static defaultProps = {
    onSave: () => {}
  };

  state = {
    langSelected: "English",
    errorType: "Only read"
  };

  render() {
    return (
      <Box pad={{ vertical: "medium", horizontal: "large" }}>
        <Heading level="3">Create new error</Heading>
        <Form
          onSubmit={e =>
            this.props.onSave({
              ...e.value,
              language: this.state.langSelected.toLowerCase(),
              type: this.state.errorType === "Only read" ? 0 : 1
            })
          }
        >
          <Box direction="row">
            <Box direction="column" margin={{ right: "large" }}>
              <Box direction="column" margin={{ vertical: "small" }}>
                <Box margin={{ horizontal: "small", vertical: "small" }}>
                  <Text>Error Type</Text>
                </Box>

                <Select
                  options={["Only read", "User Interactive"]}
                  value={this.state.errorType}
                  onChange={({ option }) =>
                    this.setState({ errorType: option })
                  }
                />
              </Box>
              <FormField label="Explain" name="explain" required />
              <FormField label="Where" name="where" />
            </Box>
            <Box direction="column">
              <Box direction="column" margin={{ vertical: "small" }}>
                <Box margin={{ horizontal: "small", vertical: "small" }}>
                  <Text>Language Message</Text>
                </Box>

                <Select
                  options={["English", "Spanish", "Chinese"]}
                  value={this.state.langSelected}
                  onChange={({ option }) =>
                    this.setState({ langSelected: option })
                  }
                />
              </Box>

              <FormField label="Message" name="message" required />
            </Box>
          </Box>

          <Button type="submit" primary label="Submit" />
        </Form>
      </Box>
    );
  }
}
