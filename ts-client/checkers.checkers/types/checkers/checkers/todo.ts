/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "checkers.checkers";

export interface Todo {
  index: string;
  creator: string;
  title: string;
  text: string;
}

function createBaseTodo(): Todo {
  return { index: "", creator: "", title: "", text: "" };
}

export const Todo = {
  encode(message: Todo, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.creator !== "") {
      writer.uint32(18).string(message.creator);
    }
    if (message.title !== "") {
      writer.uint32(26).string(message.title);
    }
    if (message.text !== "") {
      writer.uint32(34).string(message.text);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Todo {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTodo();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.creator = reader.string();
          break;
        case 3:
          message.title = reader.string();
          break;
        case 4:
          message.text = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Todo {
    return {
      index: isSet(object.index) ? String(object.index) : "",
      creator: isSet(object.creator) ? String(object.creator) : "",
      title: isSet(object.title) ? String(object.title) : "",
      text: isSet(object.text) ? String(object.text) : "",
    };
  },

  toJSON(message: Todo): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.creator !== undefined && (obj.creator = message.creator);
    message.title !== undefined && (obj.title = message.title);
    message.text !== undefined && (obj.text = message.text);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Todo>, I>>(object: I): Todo {
    const message = createBaseTodo();
    message.index = object.index ?? "";
    message.creator = object.creator ?? "";
    message.title = object.title ?? "";
    message.text = object.text ?? "";
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
