import WebsocketMessageType from "./message-type";

export default class WebsocketMessage {
    private messageType: WebsocketMessageType;
    private content: any;

    public constructor(messageType: WebsocketMessageType, content: any) {
        this.messageType = messageType;
        this.content = content;
    }
}