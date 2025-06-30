import React, { useState, useEffect } from 'react';
import styled from 'styled-components';
import { Light as SyntaxHighlighter } from 'react-syntax-highlighter';
import { atomOneDark } from 'react-syntax-highlighter/dist/esm/styles/hljs';

const Container = styled.div`
  max-width: 1200px;
  margin: 0 auto;
  padding: 6rem 2rem 4rem;
  min-height: 100vh;
`;

const Header = styled.div`
  text-align: center;
  margin-bottom: 3rem;
`;

const Title = styled.h1`
  font-size: 3rem;
  margin-bottom: 1rem;
  background: ${({ theme }) => theme.gradients.primary};
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
`;

const Subtitle = styled.p`
  font-size: 1.3rem;
  color: ${({ theme }) => theme.colors.text.secondary};
  margin-bottom: 2rem;
`;

const DownloadSection = styled.section`
  margin-bottom: 3rem;
`;

const SectionTitle = styled.h2`
  color: ${({ theme }) => theme.colors.primary};
  margin-bottom: 2rem;
  font-size: 2rem;
`;

const DownloadCard = styled.div`
  background: ${({ theme }) => theme.colors.background.secondary};
  border: 1px solid ${({ theme }) => theme.colors.border};
  border-radius: 15px;
  padding: 2rem;
  margin-bottom: 2rem;
  border: 2px solid ${({ theme }) => theme.colors.primary};
`;

const ControlsRow = styled.div`
  display: flex;
  gap: 2rem;
  margin-bottom: 2rem;
  flex-wrap: wrap;

  @media (max-width: ${({ theme }) => theme.breakpoints.mobile}) {
    flex-direction: column;
    gap: 1rem;
  }
`;

const ControlGroup = styled.div`
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
  flex: 1;
  min-width: 200px;
`;

const Label = styled.label`
  font-weight: 600;
  color: ${({ theme }) => theme.colors.text.primary};
`;

const Select = styled.select`
  padding: 0.8rem;
  border-radius: 8px;
  border: 1px solid ${({ theme }) => theme.colors.border};
  background: ${({ theme }) => theme.colors.background.tertiary};
  color: ${({ theme }) => theme.colors.text.primary};
  font-size: 1rem;
  
  &:focus {
    outline: none;
    border-color: ${({ theme }) => theme.colors.primary};
  }
`;

const DownloadButton = styled.button<{ disabled?: boolean }>`
  background: ${({ theme, disabled }) => disabled ? theme.colors.border : theme.colors.primary};
  color: white;
  padding: 1rem 2rem;
  border: none;
  border-radius: 8px;
  font-size: 1.1rem;
  font-weight: 600;
  cursor: ${({ disabled }) => disabled ? 'not-allowed' : 'pointer'};
  transition: all ${({ theme }) => theme.transitions.normal};
  min-width: 150px;

  &:hover:not(:disabled) {
    background: ${({ theme }) => theme.colors.text.accent};
    transform: translateY(-1px);
  }
`;

const InfoBox = styled.div`
  background: rgba(0, 204, 153, 0.1);
  border: 1px solid ${({ theme }) => theme.colors.primary};
  border-radius: 8px;
  padding: 1rem;
  margin: 1rem 0;
`;

const WarningBox = styled.div`
  background: rgba(255, 204, 0, 0.1);
  border: 1px solid ${({ theme }) => theme.colors.warning};
  border-radius: 8px;
  padding: 1rem;
  margin: 1rem 0;
`;

const InstallInstructions = styled.div`
  margin-top: 2rem;
`;

const VersionGrid = styled.div`
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 1.5rem;
  margin-top: 2rem;
`;

const VersionCard = styled.div`
  background: ${({ theme }) => theme.colors.background.tertiary};
  border: 1px solid ${({ theme }) => theme.colors.border};
  border-radius: 10px;
  padding: 1.5rem;
  transition: all ${({ theme }) => theme.transitions.normal};

  &:hover {
    border-color: ${({ theme }) => theme.colors.primary};
    transform: translateY(-2px);
  }
`;

const VersionTag = styled.div<{ latest?: boolean }>`
  display: inline-block;
  background: ${({ theme, latest }) => latest ? theme.colors.primary : theme.colors.border};
  color: ${({ theme, latest }) => latest ? 'white' : theme.colors.text.primary};
  padding: 0.3rem 0.8rem;
  border-radius: 12px;
  font-size: 0.9rem;
  font-weight: 600;
  margin-bottom: 1rem;
`;

const VersionDownloadBtn = styled.a`
  display: inline-block;
  background: ${({ theme }) => theme.colors.background.secondary};
  color: ${({ theme }) => theme.colors.primary};
  padding: 0.5rem 1rem;
  border-radius: 6px;
  text-decoration: none;
  border: 1px solid ${({ theme }) => theme.colors.border};
  transition: all ${({ theme }) => theme.transitions.normal};
  font-size: 0.9rem;

  &:hover {
    background: ${({ theme }) => theme.colors.primary};
    color: white;
  }
`;

const DetectedOS = styled.div`
  background: ${({ theme }) => theme.colors.background.tertiary};
  border: 1px solid ${({ theme }) => theme.colors.border};
  border-radius: 8px;
  padding: 1rem;
  margin-bottom: 1rem;
  text-align: center;
`;

interface VersionInfo {
  version: string;
  isLatest: boolean;
  releaseDate: string;
  changelog?: string;
}

const Download: React.FC = () => {
  const [selectedVersion, setSelectedVersion] = useState('v0.1.6');
  const [selectedOS, setSelectedOS] = useState('');
  const [detectedOS, setDetectedOS] = useState('');
  const [downloadUrl, setDownloadUrl] = useState('');
  const [installCommand, setInstallCommand] = useState('');

  const versions: VersionInfo[] = [
    { version: 'v0.1.6', isLatest: true, releaseDate: '2024-01-15', changelog: 'Latest features and bug fixes' },
    { version: 'v0.1.5', isLatest: false, releaseDate: '2024-01-10', changelog: 'Stable release with core features' },
    { version: 'v0.1.4', isLatest: false, releaseDate: '2024-01-05', changelog: 'Bug fixes and improvements' },
    { version: 'v0.1.3', isLatest: false, releaseDate: '2024-01-01', changelog: 'Performance improvements' },
    { version: 'v0.1.2', isLatest: false, releaseDate: '2023-12-25', changelog: 'Core functionality' },
    { version: 'v0.1.1', isLatest: false, releaseDate: '2023-12-20', changelog: 'Initial improvements' },
    { version: 'v0.1.0', isLatest: false, releaseDate: '2023-12-15', changelog: 'Initial release' },
  ];

  const osOptions = [
    { value: 'linux', label: 'Linux (64-bit)', filename: 'carrion_linux_amd64.tar.gz', arch: 'amd64' },
    { value: 'darwin', label: 'macOS (Intel & M1/M2)', filename: 'carrion-src.tar.gz', arch: 'universal' },
    { value: 'windows', label: 'Windows (64-bit)', filename: 'carrion_windows_amd64.zip', arch: 'amd64' },
    { value: 'source-zip', label: 'Source Code (.zip)', filename: '', arch: '' },
    { value: 'source-tar', label: 'Source Code (.tar.gz)', filename: '', arch: '' },
  ];

  useEffect(() => {
    // Detect user's OS
    const userAgent = navigator.userAgent.toLowerCase();
    let detected = '';
    
    if (userAgent.includes('win')) {
      detected = 'windows';
    } else if (userAgent.includes('mac')) {
      detected = 'darwin';
    } else if (userAgent.includes('linux')) {
      detected = 'linux';
    } else {
      detected = 'linux'; // Default fallback
    }
    
    setDetectedOS(detected);
    setSelectedOS(detected);
  }, []);

  useEffect(() => {
    // Update download URL when version or OS changes
    let url = '';
    
    if (selectedOS === 'source-zip') {
      url = `https://github.com/javanhut/TheCarrionLanguage/archive/refs/tags/${selectedVersion}.zip`;
    } else if (selectedOS === 'source-tar') {
      url = `https://github.com/javanhut/TheCarrionLanguage/archive/refs/tags/${selectedVersion}.tar.gz`;
    } else {
      const osOption = osOptions.find(os => os.value === selectedOS);
      if (osOption && osOption.filename) {
        url = `https://github.com/javanhut/TheCarrionLanguage/releases/download/${selectedVersion}/${osOption.filename}`;
      }
    }
    
    setDownloadUrl(url);
    
    // Update install command
    updateInstallCommand();
  // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [selectedVersion, selectedOS]);

  const updateInstallCommand = () => {
    const osOption = osOptions.find(os => os.value === selectedOS);
    if (!osOption) return;

    if (selectedOS === 'source-zip' || selectedOS === 'source-tar') {
      const fileType = selectedOS === 'source-zip' ? 'zip' : 'tar.gz';
      const extractCmd = selectedOS === 'source-zip' ? 'unzip' : 'tar -xzf';
      setInstallCommand(`# Download and build from source
curl -L "${downloadUrl}" -o carrion-${selectedVersion}.${fileType}
${extractCmd} carrion-${selectedVersion}.${fileType}
cd TheCarrionLanguage-${selectedVersion.replace('v', '')}
cd src && go build -o carrion
sudo cp carrion /usr/local/bin/

# Test installation
carrion
# In the REPL, type: version()`);
    } else if (selectedOS === 'windows') {
      setInstallCommand(`# Download and extract
curl -L "${downloadUrl}" -o carrion_windows_amd64.zip
# Extract the zip file to a folder (e.g., C:\\carrion)
# Add the folder to your PATH environment variable
# Open Command Prompt or PowerShell and test:
carrion
# In the REPL, type: version()`);
    } else if (selectedOS === 'linux') {
      setInstallCommand(`# Download and install
curl -L "${downloadUrl}" -o carrion_linux_amd64.tar.gz
tar -xzf carrion_linux_amd64.tar.gz
sudo cp carrion /usr/local/bin/
chmod +x /usr/local/bin/carrion

# Test installation
carrion
# In the REPL, type: version()`);
    } else if (selectedOS === 'darwin') {
      setInstallCommand(`# Download and install
curl -L "${downloadUrl}" -o carrion-src.tar.gz
tar -xzf carrion-src.tar.gz
sudo cp carrion /usr/local/bin/
chmod +x /usr/local/bin/carrion

# Test installation
carrion
# In the REPL, type: version()`);
    }
  };

  const handleDownload = () => {
    if (downloadUrl) {
      window.open(downloadUrl, '_blank');
    }
  };

  const getOSLabel = (osValue: string) => {
    return osOptions.find(os => os.value === osValue)?.label || osValue;
  };

  return (
    <Container>
      <Header>
        <Title>Download Carrion</Title>
        <Subtitle>Get the latest version of the Carrion programming language</Subtitle>
      </Header>

      <DownloadSection>
        <SectionTitle>Quick Download</SectionTitle>
        
        {detectedOS && (
          <DetectedOS>
            🖥️ Detected OS: <strong>{getOSLabel(detectedOS)}</strong>
          </DetectedOS>
        )}

        <DownloadCard>
          <ControlsRow>
            <ControlGroup>
              <Label>Version:</Label>
              <Select 
                value={selectedVersion} 
                onChange={(e) => setSelectedVersion(e.target.value)}
              >
                {versions.map(version => (
                  <option key={version.version} value={version.version}>
                    {version.version} {version.isLatest ? '(Latest)' : ''}
                  </option>
                ))}
              </Select>
            </ControlGroup>

            <ControlGroup>
              <Label>Operating System:</Label>
              <Select 
                value={selectedOS} 
                onChange={(e) => setSelectedOS(e.target.value)}
              >
                {osOptions.map(os => (
                  <option key={os.value} value={os.value}>
                    {os.label}
                  </option>
                ))}
              </Select>
            </ControlGroup>

            <ControlGroup>
              <Label>&nbsp;</Label>
              <DownloadButton onClick={handleDownload} disabled={!downloadUrl}>
                📥 Download {selectedVersion}
              </DownloadButton>
            </ControlGroup>
          </ControlsRow>

          {selectedOS === 'source-zip' || selectedOS === 'source-tar' ? (
            <InfoBox>
              <p><strong>Building from source requires:</strong></p>
              <ul>
                <li>Go 1.19 or later</li>
                <li>Git</li>
                <li>Basic development tools</li>
              </ul>
            </InfoBox>
          ) : selectedOS === 'windows' ? (
            <InfoBox>
              <p><strong>Download size:</strong> ~8-12 MB (zip archive)</p>
              <p><strong>Format:</strong> ZIP file containing executable</p>
              <p><strong>Requirements:</strong> No additional dependencies needed</p>
            </InfoBox>
          ) : (
            <InfoBox>
              <p><strong>Download size:</strong> ~8-12 MB (tar.gz archive)</p>
              <p><strong>Format:</strong> Compressed archive containing binary</p>
              <p><strong>Requirements:</strong> No additional dependencies needed</p>
            </InfoBox>
          )}

          <InstallInstructions>
            <h3>Installation Instructions:</h3>
            <SyntaxHighlighter 
              language={selectedOS === 'windows' ? 'powershell' : 'bash'} 
              style={atomOneDark}
              customStyle={{ margin: '1rem 0' }}
            >
              {installCommand}
            </SyntaxHighlighter>
          </InstallInstructions>

          <InstallInstructions>
            <h3>🐳 Docker Alternative:</h3>
            <SyntaxHighlighter 
              language="bash" 
              style={atomOneDark}
              customStyle={{ margin: '1rem 0' }}
            >
              {`# Pull and run ${selectedVersion === 'v0.1.6' ? 'latest' : selectedVersion.replace('v', '')} version
docker pull javanhut/carrionlanguage:${selectedVersion === 'v0.1.6' ? 'latest' : selectedVersion.replace('v', '')}
docker run -it javanhut/carrionlanguage:${selectedVersion === 'v0.1.6' ? 'latest' : selectedVersion.replace('v', '')}
# This starts the Carrion REPL automatically
# Type version() to verify the installation`}
            </SyntaxHighlighter>
          </InstallInstructions>

          {selectedOS === 'darwin' && (
            <WarningBox>
              <p><strong>macOS Users:</strong> You may need to allow the app in System Preferences → Security & Privacy if you see a security warning for unsigned binaries.</p>
            </WarningBox>
          )}
        </DownloadCard>
      </DownloadSection>

      <DownloadSection>
        <SectionTitle>All Versions</SectionTitle>
        <VersionGrid>
          {versions.map(version => (
            <VersionCard key={version.version}>
              <VersionTag latest={version.isLatest}>
                {version.version} {version.isLatest && '(Latest)'}
              </VersionTag>
              <h4>{version.version}</h4>
              <p>Released: {version.releaseDate}</p>
              <p style={{ marginBottom: '1rem', fontSize: '0.9rem' }}>{version.changelog}</p>
              <VersionDownloadBtn 
                href={`https://github.com/javanhut/TheCarrionLanguage/releases/tag/${version.version}`}
                target="_blank"
                rel="noopener noreferrer"
              >
                View Release
              </VersionDownloadBtn>
            </VersionCard>
          ))}
        </VersionGrid>
      </DownloadSection>

      <DownloadSection>
        <SectionTitle>Alternative Installation Methods</SectionTitle>
        
        <h3>📦 One-Line Installer (Linux/macOS)</h3>
        <SyntaxHighlighter language="bash" style={atomOneDark}>
{`# Using curl
curl -fsSL https://raw.githubusercontent.com/javanhut/TheCarrionLanguage/main/install/install.sh | sh

# Using wget  
wget -qO- https://raw.githubusercontent.com/javanhut/TheCarrionLanguage/main/install/install.sh | sh`}
        </SyntaxHighlighter>

        <h3>🐳 Docker</h3>
        <SyntaxHighlighter language="bash" style={atomOneDark}>
{`# Pull latest version
docker pull javanhut/carrionlanguage:latest
docker run -it javanhut/carrionlanguage:latest
# This automatically starts the Carrion REPL
# Type version() to check the version

# Pull specific version (e.g., 0.1.6)
docker pull javanhut/carrionlanguage:0.1.6
docker run -it javanhut/carrionlanguage:0.1.6

# Available tags: latest, 0.1.6, 0.1.5, 0.1.4, 0.1.3, 0.1.2, 0.1.1, 0.1.0`}
        </SyntaxHighlighter>

        <h3>📋 Package Managers (Coming Soon)</h3>
        <p>We're working on adding Carrion to popular package managers:</p>
        <ul>
          <li>Homebrew (macOS/Linux)</li>
          <li>Chocolatey (Windows)</li>
          <li>Snap (Linux)</li>
          <li>AUR (Arch Linux)</li>
        </ul>
      </DownloadSection>

      <DownloadSection>
        <SectionTitle>Verification</SectionTitle>
        <p>After installation, verify Carrion is working correctly:</p>
        <SyntaxHighlighter language="bash" style={atomOneDark}>
{`# Start the REPL to check version
carrion
# In the REPL, type: version()
# Press Ctrl+C or Ctrl+D to exit

# Run a simple program
echo 'print("Hello, Carrion!")' > hello.crl
carrion hello.crl`}
        </SyntaxHighlighter>
        
        <InfoBox>
          <p><strong>Note:</strong> The <code>carrion --version</code> flag is not yet implemented. Use the REPL command <code>version()</code> to check the installed version.</p>
        </InfoBox>
      </DownloadSection>
    </Container>
  );
};

export default Download;