let fs = require('fs')


let files = (process.argv[2] || process.env.npm_package_embedstatic_files || 'dist').split(',')
let packageName = process.argv[3] || process.env.npm_package_embedstatic_package  || 'static'
let targetFile =  process.argv[4] || process.env.npm_package_embedstatic_output || 'content.go'

var wstream = fs.createWriteStream(targetFile);

wstream.write(`package ${packageName} \n\n`)
wstream.write(`var Data = map[string][]byte {\n`)
for (let file of files) {
    addFile(file, wstream)
}
wstream.write('}')

wstream.end()


function addFile(dir, stream) {
    let stat = fs.statSync(dir)

    if (stat.isFile()) {
        stream.write(`\t"${dir}":\t[]byte("`)
        let data = fs.readFileSync(dir)

        for (let byte of data) {
            stream.write('\\x')
            if (byte < 16) {
                stream.write("0")
            }
            stream.write(byte.toString(16))
        }
        console.log("adding file: '%s', %i bytes", dir, data.length);
        stream.write(`"),\n`)
        return
    }

    let files = fs.readdirSync(dir)

    for (let file of files) {
        let uri = dir + '/' + file
        addFile(uri, stream)
    }
}
