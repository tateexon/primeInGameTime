{ stdenv, pkgs, lib }:

(pkgs.mkShell.override { stdenv = pkgs.clangStdenv; }) {
  buildInputs = with pkgs; [
    zlib
    gmp

    go_1_20
    gopls
    delve
    golangci-lint
    gotools

  ];


  # postShellHook = ''
  # '';
}