// Generated by Ignite ignite.com/cli

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient, DeliverTxResponse } from "@cosmjs/stargate";
import { EncodeObject, GeneratedType, OfflineSigner, Registry } from "@cosmjs/proto-signing";
import { msgTypes } from './registry';
import { IgniteClient } from "../client"
import { MissingWalletError } from "../helpers"
import { Api } from "./rest";
import { MsgCreateTodo } from "./types/checkers/checkers/tx";
import { MsgPlayMove } from "./types/checkers/checkers/tx";
import { MsgCreateGame } from "./types/checkers/checkers/tx";

import { Params as typeParams} from "./types"
import { StoredGame as typeStoredGame} from "./types"
import { SystemInfo as typeSystemInfo} from "./types"
import { Todo as typeTodo} from "./types"

export { MsgCreateTodo, MsgPlayMove, MsgCreateGame };

type sendMsgCreateTodoParams = {
  value: MsgCreateTodo,
  fee?: StdFee,
  memo?: string
};

type sendMsgPlayMoveParams = {
  value: MsgPlayMove,
  fee?: StdFee,
  memo?: string
};

type sendMsgCreateGameParams = {
  value: MsgCreateGame,
  fee?: StdFee,
  memo?: string
};


type msgCreateTodoParams = {
  value: MsgCreateTodo,
};

type msgPlayMoveParams = {
  value: MsgPlayMove,
};

type msgCreateGameParams = {
  value: MsgCreateGame,
};


export const registry = new Registry(msgTypes);

type Field = {
	name: string;
	type: unknown;
}
function getStructure(template) {
	const structure: {fields: Field[]} = { fields: [] }
	for (let [key, value] of Object.entries(template)) {
		let field = { name: key, type: typeof value }
		structure.fields.push(field)
	}
	return structure
}
const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
	prefix: string
	signer?: OfflineSigner
}

export const txClient = ({ signer, prefix, addr }: TxClientOptions = { addr: "http://localhost:26657", prefix: "cosmos" }) => {

  return {
		
		async sendMsgCreateTodo({ value, fee, memo }: sendMsgCreateTodoParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgCreateTodo: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgCreateTodo({ value: MsgCreateTodo.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgCreateTodo: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgPlayMove({ value, fee, memo }: sendMsgPlayMoveParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgPlayMove: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgPlayMove({ value: MsgPlayMove.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgPlayMove: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgCreateGame({ value, fee, memo }: sendMsgCreateGameParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgCreateGame: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgCreateGame({ value: MsgCreateGame.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgCreateGame: Could not broadcast Tx: '+ e.message)
			}
		},
		
		
		msgCreateTodo({ value }: msgCreateTodoParams): EncodeObject {
			try {
				return { typeUrl: "/checkers.checkers.MsgCreateTodo", value: MsgCreateTodo.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgCreateTodo: Could not create message: ' + e.message)
			}
		},
		
		msgPlayMove({ value }: msgPlayMoveParams): EncodeObject {
			try {
				return { typeUrl: "/checkers.checkers.MsgPlayMove", value: MsgPlayMove.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgPlayMove: Could not create message: ' + e.message)
			}
		},
		
		msgCreateGame({ value }: msgCreateGameParams): EncodeObject {
			try {
				return { typeUrl: "/checkers.checkers.MsgCreateGame", value: MsgCreateGame.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgCreateGame: Could not create message: ' + e.message)
			}
		},
		
	}
};

interface QueryClientOptions {
  addr: string
}

export const queryClient = ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseURL: addr });
};

class SDKModule {
	public query: ReturnType<typeof queryClient>;
	public tx: ReturnType<typeof txClient>;
	public structure: Record<string,unknown>;
	public registry: Array<[string, GeneratedType]> = [];

	constructor(client: IgniteClient) {		
	
		this.query = queryClient({ addr: client.env.apiURL });		
		this.updateTX(client);
		this.structure =  {
						Params: getStructure(typeParams.fromPartial({})),
						StoredGame: getStructure(typeStoredGame.fromPartial({})),
						SystemInfo: getStructure(typeSystemInfo.fromPartial({})),
						Todo: getStructure(typeTodo.fromPartial({})),
						
		};
		client.on('signer-changed',(signer) => {			
		 this.updateTX(client);
		})
	}
	updateTX(client: IgniteClient) {
    const methods = txClient({
        signer: client.signer,
        addr: client.env.rpcURL,
        prefix: client.env.prefix ?? "cosmos",
    })
	
    this.tx = methods;
    for (let m in methods) {
        this.tx[m] = methods[m].bind(this.tx);
    }
	}
};

const Module = (test: IgniteClient) => {
	return {
		module: {
			CheckersCheckers: new SDKModule(test)
		},
		registry: msgTypes
  }
}
export default Module;