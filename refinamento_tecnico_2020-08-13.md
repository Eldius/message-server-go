# refinamento técnico 2020-08-13 #

## finalizar adesão ##

- [Token Manager] Validar número de falhas
- [Frontend/Backend] Limitar tempo de sessão para efetuar o cadastro do cartão
- [Frontend] Tela finalizar cadastro com cartão Base Única
  - Exibir campo de cadastro de CVV
- [Frontend] Tela finalizar cadastro com novo cartão
  - Exibir dados do cartão para tokenização (com iframe)
    - Usar iframe para tokenizar cartão
  - Exibir campo de cadastro de CVV
- [Frontend] Tela de adesão com sucesso
- [Backend] Criar endpoint de cadastro de cartão e finalização da adesão
- [Token Manager] Criar endpoint para notificar falha
- [Backend] Notificar Token Manager de falha
- [Backend] Criar endpoint para o frontend notrificar o backend de falha
- [Frontend] Notificar backend de falha:
  - falha de tokenizaçào do cartão
  - falha na finalização da adesão
- [Adesão Cartão] Criar endpoint para finalizar adesão

## expirar pré cadastro ##

- Criar endpoint de expiração de pré cadastro
  - utilizar algo semelhante ao `RetentativaDeAdesaoParaFalhaDeNotificacaoActivity.expirarRetentativa`
    - status final será PRE_CADASTRO_EXPIRADO
- [Token Manager] Criar atributo de expirado sim/não no token (para validar se foi feita a expiração do fluxo de adesão)
- [Token Manager] Criar job para buscar tokens expirados ou que que atingiram maximo e ativos e expirar o pré cadastro no cartão adesão
- Token expirado:
  - Adicionar observação no evento do Call Center, para exibir o motivo `link expirado`
- Máximo de falhas
  - Adicionar observação no evento do Call Center, para exibir  `código + descrição de erro`


## obs ##

- ajustar `AssociarMeioPagamentoActivity.execute` para permitir não cancelar o pré cadastro em caso de cadastro de cartão via `cadastro-cartao`
