import React, { Component } from "react";
import { Box, Heading, Text, DataTable, Button, TextInput } from "grommet";

class App extends Component {
  state = {
    errorsData: [],
    editingRow: false
  };

  componentDidMount() {
    const ws = new WebSocket("ws://127.0.0.1:5000/ws");

    ws.onopen = () => {
      console.log("sending to something");
      ws.send("something");
    };

    ws.onmessage = data => {
      const instances = JSON.parse(data.data);
      console.log(instances);
      this.setState({ errorsData: instances });
    };
  }

  render() {
    return (
      <Box direction="column" pad="small">
        <Box direction="column">
          <Heading
            margin="none"
            level={2}
            margin="small"
            color="#1B2A32"
            responsive={true}
          >
            Welcome to Ergo Sandbox
          </Heading>
          <Text margin="small" color="#1B2A32">
            This is your workspace, where you can to see all your custom errors
            and more
          </Text>
        </Box>
        <Box>
          <DataTable
            columns={[
              {
                property: "id",
                header: <Text>ID</Text>,
                primary: true,
                render: datum => <Text size={"small"}>{datum.id}</Text>
              },
              {
                property: "code",
                header: <Text>Code</Text>,
                render: datum => (
                  <Box pad={{ vertical: "xsmall" }}>
                    {this.state.editingRow ? (
                      <TextInput
                        value={datum.code}
                        onChange={event => {
                          /* event.target.value */
                        }}
                      />
                    ) : (
                      <Text size="small">{datum.code}</Text>
                    )}
                  </Box>
                )
              },
              {
                property: "explain",
                header: <Text>Explain</Text>,
                render: datum => (
                  <Box pad={{ vertical: "xsmall" }}>
                    {this.state.editingRow ? (
                      <TextInput
                        value={datum.explain}
                        onChange={event => {
                          /* event.target.value */
                        }}
                      />
                    ) : (
                      <Text size="small">{datum.explain}</Text>
                    )}
                  </Box>
                )
              },
              {
                property: "message",
                header: <Text>User Message</Text>,
                render: datum => (
                  <Box pad={{ vertical: "xsmall" }}>
                    {this.state.editingRow ? (
                      <TextInput
                        value={datum.message}
                        onChange={event => {
                          /* event.target.value */
                        }}
                      />
                    ) : (
                      <Text size="small">{datum.message}</Text>
                    )}
                  </Box>
                )
              }
            ]}
            data={this.state.errorsData}
          />
        </Box>
      </Box>
    );
  }
}

export default App;
