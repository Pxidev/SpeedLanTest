export default class UploadTest {

    constructor() {
        this.Buffer = this.getRandomData()
    }

    // Generate Ramdom data for upload Test
    getRandomData(bufferSize = 64 * 1024) {
        const buffer = new Float32Array(new ArrayBuffer(bufferSize))

        for (let i = 0; i < buffer.length; i++) {
            buffer[i] = Math.random()
        }

        return buffer
    }

    getRandomBlob(size) {
        const data = []

        for (let i = 0; i < size / this.Buffer.byteLength; i++) {
            data.push(this.Buffer)
        }

        return new Blob(data, {
            type: "application/octet-stream"
        })

    }


}