} else if walletSigner != nil {
  // If walletSigner is set, it must be a signer for the key being used
  // for the transaction.
  if err := walletSigner.Sign(ctx, tx); err != nil {
    return err
  }
}  
