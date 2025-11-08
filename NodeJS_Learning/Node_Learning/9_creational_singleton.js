class Server {
    constructor(port) {
      this._port = port;
    }
    static init(port) {
    }
    static getInstance() {
    }
    status() {
      console.log("Server listening on port " + this._port);
    }
  }
  
  /**
   * Client calls init, and getInstance would give that instance
   * always. Singleton is used for heavy single use objects like DB
   */
  Server.init(1234);
  Server.getInstance().status();