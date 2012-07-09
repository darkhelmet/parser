package parser

func Between(left, right string, inside Parser) Parser {
    return func(r *Reader) (interface{}, *Reader, error) {
        results, rd, err := And(String(left), inside, String(right))(r)
        if err != nil {
            return results, r, err
        }
        result := results.([]interface{})[1]
        return result, rd, nil
    }
}
