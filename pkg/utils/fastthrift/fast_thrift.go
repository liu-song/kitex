/*
 * Copyright 2022 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package fastthrift

import "github.com/liu-song/kitex/pkg/remote/codec/thrift"

// FastMarshal marshals the msg to buf. The msg should be generated by Kitex tool and implement ThriftMsgFastCodec.
func FastMarshal(msg thrift.ThriftMsgFastCodec) []byte {
	buf := make([]byte, msg.BLength())
	msg.FastWriteNocopy(buf, nil)
	return buf
}

// FastUnmarshal unmarshal the buf into msg. The msg should be generated by Kitex tool and implement ThriftMsgFastCodec.
func FastUnmarshal(buf []byte, msg thrift.ThriftMsgFastCodec) error {
	_, err := msg.FastRead(buf)
	return err
}
