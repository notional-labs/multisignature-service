import {StdSignDoc, makeSignDoc, Coin,} from "@cosmjs/launchpad"
import {AminoMsg, StdFee} from "./type" 

function getKeplr(): Promise<Keplr | undefined> {
    if (window.keplr) {
        return window.keplr;
    }
    
    if (document.readyState === "complete") {
        return window.keplr;
    }
    
    return new Promise((resolve) => {
        const documentStateChange = (event: Event) => {
            if (
                event.target &&
                (event.target as Document).readyState === "complete"
            ) {
                resolve(window.keplr);
                document.removeEventListener("readystatechange", documentStateChange);
            }
        };
        
        document.addEventListener("readystatechange", documentStateChange);
    });
}

function createStdSigDoc(json_data: JSON): StdSignDoc{
    var msg : AminoMsg = {
        type :json_data.parse("type"),
        value : {
            from_address: json_data.parse("from_address"),
            to_address: json_data.parse("to_address"),
            amount: json_data.parse("amount"),
        }
    }
    var arr_amino_msg : AminoMsg[]
    arr_amino_msg.push(msg)

    var coin : Coin = {
        denom: json_data.parse("denom"),
        amount: json_data.parse('amount'),
    }
    var arr_coin : Coin[]
    arr_amino_msg.push(coin)


    var fee : StdFee = {
        amount: arr_coin,
        gas:json_data.parse("gas")
    }
    

    return makeSignDoc(
        arr_amino_msg, 
        fee, 
        json_data.parse("chainId"), 
        json_data.parse("memo"), 
        json_data.parse("accountNumber"), 
        json_data.parse("sequence")
    )
}

