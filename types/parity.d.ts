/** spell-checker: disable */
export interface IParityTxTrace {
  action: IParityCallAction | IParityCreateAction | IParitySuicideAction;
  blockHash: string;
  blockNumber: number;
  // if type == "suicide" then result is null
  // if `error` !== void(0) then `result` is `void(0)`
  result?: IParityCallResult | IParityCreateResult | null;
  subtraces: number;
  error?: string;
  traceAddress: number[];
  transactionHash: string;
  transactionPosition: number;
  type: "create" | "call" | "suicide";
}

export interface IParityCallAction {
  callType: "call" | "callcode" | "delegatecall" | "staticcall";
  from: string;
  to: string;
  value: string;
  gas: string;
  input: string;
}

export interface IParityCreateAction {
  from: string;
  value: string;
  gas: string;
  init: string;
}

export interface IParitySuicideAction {
  address: string;
  refundAddress: string;
  balance: string;
}

export interface IParityCreateResult {
  address: string;
  code: string;
  gasUsed: string;
}

export interface IParityCallResult {
  gasUsed: string;
  output: string;
}
