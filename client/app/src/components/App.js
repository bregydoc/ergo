import React, { Component } from "react";
import {
  Box,
  Heading,
  Text,
  DataTable,
  Button,
  TextInput,
  Layer
} from "grommet";
import * as Icons from "grommet-icons";
import NewErrorForm from "./NewErrorForm";

import axios from "axios";

class App extends Component {
  state = {
    errorsData: [],
    editingRow: false,
    creatingNewError: false
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

  saveNewError = data => {
    axios.post("http://127.0.0.1:5000/save-error", data).then(r => {
      console.log(r.status);
      console.log(r.data);
    });
  };

  render() {
    return (
      <div>
        {this.state.creatingNewError ? (
          <Layer
            onClickOutside={() => this.setState({ creatingNewError: false })}
          >
            <div>
              <NewErrorForm onSave={this.saveNewError} />
            </div>
          </Layer>
        ) : null}

        <Box direction="column" pad="small">
          <Box
            direction="row-responsive"
            alignContent="center"
            align="center"
            pad={{ bottom: "small" }}
          >
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
                This is your workspace, here you can to see all your custom
                errors and more
              </Text>
            </Box>
            <Box alignContent="end" align="end" pad="small" flex="grow">
              <Button
                primary
                icon={<Icons.Add />}
                label="New Error"
                onClick={() => {
                  this.setState({ creatingNewError: true });
                }}
              />
            </Box>
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
                  property: "type",
                  header: <Text>Type</Text>,
                  primary: true,
                  render: datum => <Text size={"small"}>{datum.type}</Text>
                },
                {
                  property: "code",
                  header: <Text>Code</Text>,
                  render: datum => (
                    <Box pad={{ vertical: "xsmall" }}>
                      {this.state.editingRow === datum.id ? (
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
                      {this.state.editingRow === datum.id ? (
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
                      {this.state.editingRow === datum.id ? (
                        <TextInput
                          value={datum.english_user_message}
                          onChange={event => {
                            /* event.target.value */
                          }}
                        />
                      ) : (
                        <Text size="small">{datum.english_user_message}</Text>
                      )}
                    </Box>
                  )
                },
                {
                  // property: "nil",
                  header: <Text>Actions</Text>,
                  render: datum => {
                    return this.state.editingRow !== datum.id ? (
                      <Box
                        pad={{ vertical: "xsmall" }}
                        direction="row-responsive"
                      >
                        <Button
                          icon={<Icons.Edit />}
                          label="Edit"
                          color="#1B2A32"
                          // hoverIndicator={"#1B2A32"}
                          onClick={() => {
                            if (this.state.editingRow !== datum.id) {
                              this.setState({ editingRow: datum.id });
                            } else {
                              this.setState({ editingRow: false });
                            }
                            console.log(this.state.editingRow);
                          }}
                          primary
                        />

                        <Box pad={{ left: "small" }}>
                          <Button
                            icon={<Icons.Trash />}
                            label="Delete"
                            color="#FF6764"
                            hoverIndicator={true}
                            onClick={() => {}}
                          />
                        </Box>
                      </Box>
                    ) : (
                      <Box
                        pad={{ vertical: "xsmall" }}
                        direction="row-responsive"
                      >
                        <Button
                          icon={<Icons.Checkmark />}
                          label="Save"
                          color="#01E69F"
                          hoverIndicator={true}
                          onClick={() => {}}
                          primary
                        />

                        <Box pad={{ left: "small" }}>
                          <Button
                            icon={<Icons.Close />}
                            label="Cancel"
                            color="#FF6764"
                            hoverIndicator={true}
                            onClick={() => {
                              this.setState({ editingRow: false });
                            }}
                          />
                        </Box>
                      </Box>
                    );
                  }
                }
              ]}
              data={this.state.errorsData}
            />
          </Box>
        </Box>
      </div>
    );
  }
}

export default App;
