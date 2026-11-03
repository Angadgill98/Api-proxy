package proxy

import "github.com/google/nftables"

const TableName = "api_debugger"

func CreateTable() (*nftables.Table, error) {
    conn := &nftables.Conn{}

    table := conn.AddTable(&nftables.Table{
        Family: nftables.TableFamilyIPv4,
        Name:   TableName,
    })

    if err := conn.Flush(); err != nil {
        return nil, err
    }

    return table, nil
}

func CreateBaseChain(conn *nftables.Conn, table *nftables.Table) (*nftables.Chain, error) {
    chain := conn.AddChain(&nftables.Chain{
        Name:     "output",
        Table:    table,
        Type:     nftables.ChainTypeNAT,
        Hooknum:  nftables.ChainHookOutput,
        Priority: nftables.ChainPriorityNATDest,
    })

    if err := conn.Flush(); err != nil {
        return nil, err
    }

    return chain, nil
}