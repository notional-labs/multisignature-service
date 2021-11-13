import {Coin} from "@cosmjs/launchpad"
export interface Msg {
    readonly type: string;
    readonly value: any;
}

export interface AminoMsg {
    readonly type: string;
    readonly value: any;
}
/** A high level transaction of the coin module */
export interface MsgSend extends Msg {
    readonly type: "cosmos-sdk/MsgSend";
    readonly value: {
        /** Bech32 account address */
        readonly from_address: string;
        /** Bech32 account address */
        readonly to_address: string;
        readonly amount: readonly Coin[];
    };
}

export interface StdFee {
    readonly amount: readonly Coin[];
    readonly gas: string;
}
