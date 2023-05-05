{ stdenv, pkgs, lib }:

(pkgs.mkShell.override { stdenv = pkgs.clangStdenv; }) {
  buildInputs = with pkgs; [
    zlib # for numpy
    gmp

    go_1_20
    gopls
    delve
    golangci-lint
    gotools
    # make

  ];


  # postShellHook = ''
  # '';
}