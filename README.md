# _Machinebox-Textbox_ OMG Microservice

[![Open Microservice Guide](https://img.shields.io/badge/OMG%20Enabled-👍-green.svg?)](https://microservice.guide)
[![Build Status](https://travis-ci.com/omg-services/machinebox-textbox.svg?branch=master)](https://travis-ci.com/omg-services/machinebox-textbox)
[![codecov](https://codecov.io/gh/omg-services/machinebox-textbox/branch/master/graph/badge.svg)](https://codecov.io/gh/omg-services/machinebox-textbox)


An OMG service for machinebox-textbox, it understand text with natural language processing (NLP) and analyze the sentiment and find keywords from given text.

## Direct usage in [Storyscript](https://storyscript.io/):

##### Text Analyze
```coffee
machinebox-textbox analyze text:'I can't wait for @machineboxio to release Textbox; it provides natural language processing and a whole host of other #useful things.'
{"sentences": ["List of text and analyze data"],"keywords": ["List of keywords"]}
```

Curious to [learn more](https://docs.storyscript.io/)?

✨🍰✨

## Usage with [OMG CLI](https://www.npmjs.com/package/omg)

##### Text Analyze
```shell
$ omg run analyze -a text=<TEXT/SENTENCE_TO_ANALYZE> -e ADDRESS=<SERVER_IP_ADDRESS_WITH_PORT> 
```

**Note**: The OMG CLI requires [Docker](https://docs.docker.com/install/) to be installed.

## License
[MIT License](https://github.com/omg-services/machinebox-textbox/blob/master/LICENSE).
