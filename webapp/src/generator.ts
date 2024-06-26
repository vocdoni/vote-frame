type DataType = string | number

export class CsvGenerator {
  constructor(
    private headers: string[],
    private data: DataType[][],
    public filename?: string
  ) {}

  public get url(): string {
    let binaryData = 'data:application/csv,'

    for (const index in this.headers) {
      let separator = '%2C'

      if (parseInt(index) + 1 === this.headers.length) {
        separator = '%0A'
      }

      binaryData += `${this.headers[index]}${separator}`
    }

    for (let rowIndex = 0; rowIndex < this.data.length; ++rowIndex) {
      const row = this.data[rowIndex]

      for (let fieldIndex = 0; fieldIndex < row.length; ++fieldIndex) {
        let separator = '%2C'

        if (fieldIndex + 1 === row.length) {
          separator = '%0A'
        }

        binaryData += `${row[fieldIndex]}${separator}`
      }
    }

    return binaryData
  }
}
