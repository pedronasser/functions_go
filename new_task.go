/* 
 * IronFunctions
 *
 * The open source serverless platform.
 *
 * OpenAPI spec version: 0.1.27
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package swagger

type NewTask struct {

	// Name of Docker image to use. This is optional and can be used to override the image defined at the group level.
	Image string `json:"image,omitempty"`

	// Payload for the task. This is what you pass into each task to make it do something.
	Payload string `json:"payload,omitempty"`
}
