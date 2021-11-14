import {StdSignDoc, makeSignDoc, Coin, StdSignature} from "@cosmjs/launchpad"
import {Msg, StdFee} from "./type" 
class KeplrHandler{
    static getKeplr(): Keplr {
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

    private createStdSigDoc(json_data:any): StdSignDoc{        
        var msg : Msg = {
            type :json_data.parse("type"),
            value : {
                from_address: json_data.parse("from_address"),
                to_address: json_data.parse("to_address"),
                amount: json_data.parse("amount"),
            }
        }
        var arr_amino_msg : Msg[] = [msg]
        var coin : Coin = {
            denom: json_data.parse("denom"),
            amount: json_data.parse('amount'),
        }
        var arr_coin : Coin[] = [coin]

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
    private getJson(): any{
        //TODO: REST API call
    }

    public async getSignature(): Promise<any>{
        var json_data = this.getJson()
        var keplr = this.getKeplr();
        
        let data = await keplr.signAmino(this.createStdSigDoc(json_data.parse("tx_id")))

        return {
            tx_id    : json_data.parse("tx_id"),
            signature: data.signature.signature
        }
    }

}