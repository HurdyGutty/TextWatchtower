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

export namespace reloadPoint {
	
	export class MousePoint {
	    X: number;
	    Y: number;
	
	    static createFrom(source: any = {}) {
	        return new MousePoint(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.X = source["X"];
	        this.Y = source["Y"];
	    }
	}

}

export namespace screenBox {
	
	export class ScreenBox {
	    x: number;
	    y: number;
	    w: number;
	    h: number;
	
	    static createFrom(source: any = {}) {
	        return new ScreenBox(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.x = source["x"];
	        this.y = source["y"];
	        this.w = source["w"];
	        this.h = source["h"];
	    }
	}

}

