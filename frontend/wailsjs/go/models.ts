export namespace instruct {
	
	export class Instruct {
	    message: string;
	    state: string;
	
	    static createFrom(source: any = {}) {
	        return new Instruct(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.message = source["message"];
	        this.state = source["state"];
	    }
	}

}

