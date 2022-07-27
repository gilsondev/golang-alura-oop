package contas

type Conta interface {
	Sacar(valorDoSaque float64) string
	Depositar(valorDoDeposito float64) (string, float64)
	Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool
	Saldo() float64
}
