// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { createLogic } from 'redux-logic'

import api from '@oauth/api'

import * as init from '@oauth/store/actions/init'
import * as user from '@oauth/store/actions/user'

const oauthInitLogic = createLogic({
  type: init.INITIALIZE,
  async process({ getState, action }, dispatch, done) {
    dispatch(user.getUserMe())

    try {
      try {
        const result = await api.oauth.me()

        dispatch(user.getUserMeSuccess(result.data.user))
      } catch (error) {
        const userError = error.data ? error.data : error
        dispatch(user.getUserMeFailure(userError))
      }

      dispatch(init.initializeSuccess())

      // eslint-disable-next-line no-console
      console.log('Oauth initialization successful!')
    } catch (error) {
      const initError = error.data ? error.data : error
      dispatch(init.initializeFailure(initError))
    }

    done()
  },
})

export default oauthInitLogic
