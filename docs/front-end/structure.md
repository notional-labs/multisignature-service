# FRONT - END STRUCTURE
1. cosmos-logic:
	 * class KeplrHandler:
	   * static getKeplr(): new KeplrHandler() || null
	   	 (connect to user wallet. If success, new KeplrHandler() else null)
	   	 
	   * static checkKeplr(): boolean
	   	 (check if browser has keplr installed)
	   
		 * getSignature(tx: Object): Object || null
		   (prompting user's keplr to get signature) 
		   
		 * verifyTx(tx: Object): boolean
		   (verify if a transaction is valid)
		   
	 * class TxHandler:
	 	 * (STILL DETERMINING PARAMS) static generateTx(): Object
	 	   (generate transaction)
	 	   
2. the rest of front - end:
   * exchanging data with back - end
   * provide basic UI to user
